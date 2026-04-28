# pkg/rule

## 변경이력

- 2026-04-28: 4 원칙 준수 형식으로 개정

## 역할

Toulmin defeats graph 기반 warrant 함수 라이브러리. `pkg/validate` 하위 검증 규칙 중 약 1/4 (35 개) 가 본 패키지의 warrant + defeater 를 사용하며 나머지는 `*Ground` 직접 조회 plain function. **규칙 wiring 은 `pkg/validate/<domain>/run.go`**, **규칙 카탈로그는 [`rulebook.md`](../../rulebook.md)**.

> 관련: [`pkg/yongol`](../yongol) (Fullstack) · [`pkg/ground`](../ground) (Build adapter) · [`pkg/validate`](../validate) (wiring) · [`pkg/diagnostic`](../diagnostic).

## 공개 함수 (warrant + defeater + util)

| 함수 | 시그니처 | 설명 |
|---|---|---|
| `FieldRequired` | `(toulmin.Context, toulmin.Specs) (bool, any)` | S-1~S-24 (22 규칙). HTTP 전용 필드 요구 (claim: `map[string]bool`). `IsSubscribe` defeater |
| `VarDeclared` | 동일 | S-27~S-30, XSS-47 (5 규칙). `Ground.Vars` 조회 (claim: `string`). `IsImplicitVar` defeater (예약어 `currentUser`/`request`/`query`/`message` 면제) |
| `CoverageCheck` | 동일 | XOD-10, XSO-20, XSD-55, XPN-54, XPN-64 (5 규칙). `Ground.Lookup[LookupKey]` 조회. defeater 연계 (`IsSensitiveCol`, `IsArchived`+`IsPkgModel` 등) |
| `IsPkgModel` | `(...) (bool, any)` | defeater. `Flags["pkgModel"]`. 면제: XSD-55, XDS-12 |
| `IsArchived` | 동일 | `Flags["archived"]`. 면제: XSD-55 (DDL `@archived`) |
| `IsSensitiveCol` | 동일 | `Flags["sensitive"]`. 면제: XOD-10 (DDL `-- @sensitive`) |
| `IsNoSensitive` | 동일 | `Flags["nosensitive"]`. 면제: XDD-61 |
| `IsSubscribe` | 동일 | `Flags["subscribe"]`. 면제: FieldRequired (S-1~24) — @subscribe 함수 |
| `IsImplicitVar` | 동일 | `Flags["implicit.<name>"]`. 면제: VarDeclared (S-27~30, XSS-47) |
| `IsCustomTS` | 동일 | `Flags["customTS.<name>"]`. 면제: TM-8 (`<page>.custom.ts` export) |
| `GoInitialism` / `IsUpper` / `SplitPascal` / `StringSet` | (util) | 네이밍/문자열 헬퍼 |

## 공개 타입

| 타입 | 설명 |
|---|---|
| `Ground` | `{Lookup, Types, Pairs, Schemas map[string]...; Config map[string]bool; Vars, Flags StringSet}` — `pkg/ground.Build` 가 채움 |
| `Evidence` | `{Rule, Level, Ref, Message string}` — 위반 결과 (Level: `ERROR`/`WARNING`) |
| `BaseSpec` | `{Rule, Level, Message string}` — 모든 spec embed. `toulmin.Spec` 구현 (`SpecName()`, `Validate()`) |
| `FieldRequiredSpec` | `BaseSpec + {SeqType, Field string; Required bool}` |
| `VarDeclaredSpec` | `BaseSpec` (claim: 변수명 string) |
| `CoverageCheckSpec` | `BaseSpec + {LookupKey string}` |
| `StringSet` | `map[string]bool` 별칭 + helper |

## Toulmin 정당화 조건 (warrant 추가 기준)

다음 둘 중 하나 이상일 때만 Toulmin 그래프 사용:

1. defeater 가 실제로 판정을 뒤집는다 (`IsImplicitVar`, `IsSubscribe`, `IsCustomTS`, `IsSensitiveCol`, `IsArchived`, `IsPkgModel`, `IsNoSensitive`).
2. 스킵/예외 조건이 다수이며 확장 가능성 (예: XSS-11, XDS-12 — 4+ 스킵 체인).

해당 안 되고 "Ground.Lookup 집합 조회"로 끝나면 plain `func(g *Ground, ...)` + if-else.

## 시그니처 규약

```go
func RuleName(ctx toulmin.Context, specs toulmin.Specs) (bool, any)
```

`ctx.Get("claim")` (검증 대상), `ctx.Get("ground")` (`*Ground`), `specs[0].(*XxxSpec)`. 반환 `(true, *Evidence)` = 위반, `(false, nil)` = 통과. ctx 주입은 `pkg/validate` 호출자 책임.

## 커스텀 warrant (`pkg/rule` 공통 함수 미사용)

`pkg/validate/{folder}` 내 자체 warrant + Toulmin graph 등록. 공통화 어려움:

- `XSS-11` PluralResultType — primitive/call/package 스킵 체인
- `XDS-12` ResultNoDDLTable — `IsPkgModel` + sqlc row type + seq type 스킵
- `XDD-61` SensitiveNoAnnotation — `IsNoSensitive` 면제
- `TM-8` BindNotFound — `IsCustomTS` 면제

## 제거된 공통 warrant (2026-04)

과거 `RefExists`, `PairMatch`, `TypeMatch`, `SchemaMatch`, `ConfigRequired`, `ForbiddenRef`, `NameFormat` 7 종 + `SchemaEvidence`/`TypeClaim` 보조 타입 export 했음. 검토 결과 모두 defeater 없이 단순 Ground 조회로 환원 → `pkg/validate` 각 폴더 plain `g.Lookup[key][name]` 으로 대체 후 삭제. 재도입 시 (1) 진짜 defeater 존재 검토, (2) `bak/pkg/rule/` 참조, (3) PR 에 defeater 가 판정을 뒤집는지 입증.

## 네이밍 규약

- warrant/defeater 함수: PascalCase (`FieldRequired`, `IsPkgModel`).
- spec 타입: `<RuleName>Spec` (BaseSpec embed).
- LookupKey/ConfigKey 는 `pkg/ground` populator 가 채우는 키와 정확히 일치 (silent-pass 방지).

## catalog/

`pkg/rule/catalog/` 에 규칙 카탈로그 보조 메타. 카탈로그 진실 원본은 저장소 루트 `rulebook.md`.
