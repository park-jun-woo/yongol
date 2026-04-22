# pkg/rule

> **관계**: 본 패키지는 Toulmin defeater 그래프 기반 warrant 함수 라이브러리다.
> - **규칙 wiring / 오케스트레이션** 은 [`pkg/validate/`](../validate) 참조.
> - `pkg/rule` 이 제공하는 것은 **warrant 로직** (예: `VarDeclared`, `CoverageCheck`, `FieldRequired`) + **predicate** (`IsArchived` 등) + **Spec 타입**.
> - 실제 "어떤 검증이 어떤 SSOT 에서 돌아가는가" 는 `pkg/validate/<domain>/run.go` 의 호출 순서에서 확인.
> - **규칙 카탈로그** (rule ID → 설명): 저장소 루트 [`rulebook.md`](../../rulebook.md).

Toulmin defeats graph 규칙 라이브러리. `pkg/validate` 하위 검증 규칙 중 **대략 1/4 수준** 이 본 패키지의 warrant/defeater를 사용한다. 나머지 규칙은 `*Ground`를 직접 조회하는 plain function으로 구현된다.

> 관련 패키지: [`pkg/yongol`](../yongol) (Fullstack) · [`pkg/ground`](../ground) (Fullstack → Ground 어댑터) · [`pkg/validate`](../validate) (오케스트레이션 + 규칙 wiring) · [`pkg/diagnostic`](../diagnostic) (공용 진단 타입)

## Toulmin이 정당화되는 조건

두 조건 중 하나 이상을 충족할 때만 Toulmin 그래프를 사용한다:

1. **defeater가 실제로 판정을 뒤집는다** — `IsImplicitVar`, `IsSubscribe`, `IsCustomTS`, `IsSensitiveCol`, `IsArchived`, `IsPkgModel`, `IsNoSensitive` 가 예외 케이스를 면제
2. **스킵/예외 조건이 다수이며 추후 확장 가능성이 있다** — 예: `XSS-11`, `XDS-12` (4+ 스킵 체인)

위 조건에 해당하지 않고 "Ground.Lookup 집합 조회"로 끝나면 **plain `func(g *Ground, ...)` + if-else** 로 구현한다.

## 시그니처 규약

```go
func RuleName(ctx toulmin.Context, specs toulmin.Specs) (bool, any)
```

- `ctx.Get("claim")` — 검증 대상 (이름, 변수, 필드 집합 등)
- `ctx.Get("ground")` — `*Ground` 공유 조회 컨텍스트
- `specs[0].(*XxxSpec)` — 규칙별 판정 기준 struct
- **반환**: `(true, *Evidence)` = 위반 발생, `(false, nil)` = 통과

claim/ground 는 graph 구성 시 `pkg/validate` 호출자가 ctx에 주입한다.

## 타입

### Ground

`pkg/ground.Build(fs *yongol.Fullstack)`가 구축. populator가 채우는 키 목록은 `pkg/ground/README.md` 참조.

```go
type Ground struct {
    Lookup  map[string]StringSet // "target.kind" -> set of names
    Types   map[string]string    // "target.kind.name" -> type string
    Pairs   map[string]StringSet // "target.pairKind" -> set of "key:value"
    Config  map[string]bool      // config key -> present
    Vars    StringSet            // declared variable names
    Flags   StringSet            // flags for defeaters
    Schemas map[string][]string  // "target.schema" -> ordered field list
}
```

### Evidence

```go
type Evidence struct {
    Rule    string // "S-1", "XSS-11" 등
    Level   string // "ERROR" 또는 "WARNING"
    Ref     string
    Message string
}
```

### BaseSpec

모든 `*XxxSpec` 타입이 embed 하는 공통 필드. toulmin.Spec 인터페이스를 만족 (`SpecName()`, `Validate()`).

```go
type BaseSpec struct {
    Rule    string
    Level   string
    Message string
}
```

## Warrant — 현재 Toulmin으로 wiring 된 3종

### FieldRequired (22 규칙)

`pkg/validate/ssac` S-1~S-24. `IsSubscribe` defeater가 HTTP 전용 필드 요구를 @subscribe 함수에서 면제.

```go
type FieldRequiredSpec struct {
    BaseSpec
    SeqType  string // "@get", "@post", "@put", "@delete", "@empty", "@state", "@auth", "@call", "@publish"
    Field    string // "Model", "Result", "Inputs", "Target", "Message" 등
    Required bool   // true = 있어야 함, false = 없어야 함
}
```

claim: `map[string]bool` (필드명 → 존재 여부).

### VarDeclared (5 규칙)

`pkg/validate/ssac` S-27~S-30 + XSS-47. `IsImplicitVar` defeater가 예약어(`currentUser`/`request`/`query`/`message`) 면제.

```go
type VarDeclaredSpec struct {
    BaseSpec
}
```

claim: `string` (변수명). `Ground.Vars` 조회.

### CoverageCheck (5 규칙)

defeater 연계로 사용: XOD-10 (`IsSensitiveCol`), XSO-20 (Wrapper 스킵 — 향후 defeater 전환 가능성), XSD-55 (`IsArchived`+`IsPkgModel`), XPN-54/XPN-64 (middleware/response 참조 검토 중).

```go
type CoverageCheckSpec struct {
    BaseSpec
    LookupKey string
}
```

claim: `string` (정의된 항목). `Ground.Lookup[LookupKey]` 조회.

## Defeater

모두 `Ground.Flags` 기반. 호출자가 평가 전에 해당 플래그를 설정.

| 함수 | spec | Flags 키 | 면제 대상 규칙 |
|---|---|---|---|
| `IsPkgModel` | nil | `pkgModel` | XSD-55, XDS-12 — `pkg/<pkg>/` 내장 모델 |
| `IsArchived` | nil | `archived` | XSD-55 — DDL `@archived` |
| `IsSensitiveCol` | nil | `sensitive` | XOD-10 — DDL `-- @sensitive` |
| `IsNoSensitive` | nil | `nosensitive` | XDD-61 — DDL `-- @nosensitive` |
| `IsSubscribe` | nil | `subscribe` | FieldRequired (S-1~24 HTTP 전용) — @subscribe 함수 |
| `IsImplicitVar` | nil | `implicit.<name>` | VarDeclared (S-27~30, XSS-47) — 예약어 |
| `IsCustomTS` | nil | `customTS.<name>` | TM-8 — `<page>.custom.ts` export |

## 커스텀 warrant (`pkg/rule` 공통 함수 미사용)

아래 규칙은 `pkg/validate/{folder}` 내에 자체 warrant 함수를 두고 Toulmin graph에 defeater와 함께 등록한다. 공통 warrant로 표현하기엔 고유 로직이 많음:

- `XSS-11` PluralResultType — primitive/call/package 스킵 체인
- `XDS-12` ResultNoDDLTable — `IsPkgModel`+sqlc row type+seq type 스킵
- `XDD-61` SensitiveNoAnnotation — `IsNoSensitive` 면제
- `TM-8` BindNotFound — `IsCustomTS` 면제

## 제거된 공통 warrant (2026-04)

과거 `pkg/rule`은 `RefExists`, `PairMatch`, `TypeMatch`, `SchemaMatch`, `ConfigRequired`, `ForbiddenRef`, `NameFormat` 7종 + `SchemaEvidence`/`TypeClaim` 보조 타입을 export 했다. 검토 결과 이들 모두 **defeater 없이 단순 Ground 조회**로 환원되어, `pkg/validate` 각 폴더에서 plain `g.Lookup[key][name]` 체크로 대체 가능함이 확인돼 **삭제**.

향후 동일 패턴을 다시 필요로 하는 경우:
1. 먼저 **진짜 defeater가 있는지** 검토 — 없으면 plain 함수 유지
2. defeater가 있으면 `bak/pkg/rule/`에서 구현 참조해 재도입
3. 신규 warrant 추가 시 해당 defeater가 **실제로 판정을 뒤집는지** PR 설명에 입증할 것

## 네이밍 규약

- warrant/defeater 함수: `PascalCase` — `FieldRequired`, `IsPkgModel`
- spec 타입: `<RuleName>Spec` — `FieldRequiredSpec`, `CoverageCheckSpec`
- 모든 Spec은 `BaseSpec` embed
- LookupKey/ConfigKey 네임스페이스는 `pkg/ground` populator가 채우는 것과 일치해야 함 (silent-pass 방지)
