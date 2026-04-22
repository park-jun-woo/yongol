# pkg/validate

> **관계**: 본 패키지는 전 validate 규칙의 오케스트레이션·wiring 을 담당한다.
> - **Toulmin defeater 그래프 warrant 함수** 는 [`pkg/rule/`](../rule) 참조.
> - 본 패키지의 각 `<domain>/run.go` 가 `pkg/rule` 의 warrant + predicate 를 사용해 최종 diag 을 emit.
> - **규칙 카탈로그** 는 저장소 루트 [`rulebook.md`](../../rulebook.md).
> - **구 `pkg/crosscheck/`** — 2026-04 통합. 교차 검증은 본 패키지 `<sourceA>_<sourceB>/` 서브 디렉토리에 위치 (예: `openapi_ssac`, `ssac_rego`).

yongol 의 단일·교차 정합성 검증 통합 패키지. `pkg/rule` 공통 함수 + 고유 함수로 구성.

> **정합성 기준**: 본 문서는 `internal/ssac/validator`, `internal/stml/validator`, `internal/crosscheck` 등 테스트로 확정된 internal 구현의 거동을 Toulmin defeats 그래프 위에서 재구성한 것이다.

## 패키지 구조

### 단일 SSOT 검증 (per-SSOT)

| 패키지 | parser 대응 | 규칙 ID 접두사 | 설명 |
|--------|------------|---------------|------|
| [`manifest/`](manifest/) | `parser/manifest/` | `C-` | manifest.yaml 로드 검증 |
| [`ddl/`](ddl/) | `parser/ddl/` | `D-`, `XDD-` | sqlc 중복, NOT NULL, 센티널, 민감 패턴 컬럼 |
| [`openapi/`](openapi/) | kin-openapi 직접 | `O-`, `XOO-` | path 파라미터 충돌, password minLength, email format |
| [`ssac/`](ssac/) | `parser/ssac/` | `S-`, `XSS-` | 필수 필드, 변수 흐름, 모델 검증, @subscribe 제약, pub/sub 쌍 |
| [`stml/`](stml/) | `parser/stml/` | `TM-` | fetch/action 바인딩, 파라미터, 컴포넌트 참조 |
| [`statemachine/`](statemachine/) | `parser/statemachine/` | `ST-` | Mermaid stateDiagram 파싱 검증 |
| [`rego/`](rego/) | `parser/rego/` | `P-`, `XPP-` | Rego 파싱, @ownership 누락 |
| [`func/`](func/) | `parser/funcspec/` | `F-`, `XFF-` | built-in 패키지명 충돌, func 본체 TODO, 금지 import |
| [`hurl/`](hurl/) | `parser/hurl/` | `H-` | `.feature` 파일 deprecated |

### 교차 SSOT 검증 (SSOT pair)

| 폴더 | 규칙 ID 접두사 | 관련 SSOT |
|------|---------------|-----------|
| [`openapi_ddl/`](openapi_ddl/) | `XDO-`, `XOD-` | OpenAPI ↔ DDL |
| [`openapi_ssac/`](openapi_ssac/) | `XOS-`, `XSO-` | OpenAPI ↔ SSaC |
| [`openapi_hurl/`](openapi_hurl/) | `XOH-` | OpenAPI ↔ Hurl |
| [`openapi_manifest/`](openapi_manifest/) | `XNO-`, `XON-` | OpenAPI ↔ Manifest |
| [`ssac_ddl/`](ssac_ddl/) | `XSD-`, `XDS-` | SSaC ↔ DDL |
| [`ssac_statemachine/`](ssac_statemachine/) | `XMS-`, `XSM-` | SSaC ↔ StateMachine |
| [`ssac_func/`](ssac_func/) | `XFS-`, `XSF-` | SSaC ↔ Func |
| [`ssac_manifest/`](ssac_manifest/) | `XNS-` | SSaC ↔ Manifest |
| [`ssac_rego/`](ssac_rego/) | `XPS-`, `XSP-` | SSaC ↔ Rego |
| [`ssac_authz/`](ssac_authz/) | `XAS-` | SSaC ↔ Authz |
| [`ssac_sqlc/`](ssac_sqlc/) | `XQS-` | SSaC ↔ sqlc |
| [`ddl_statemachine/`](ddl_statemachine/) | `XDM-` | DDL ↔ StateMachine |
| [`ddl_rego/`](ddl_rego/) | `XDP-` | DDL ↔ Rego |
| [`rego_manifest/`](rego_manifest/) | `XNP-`, `XPN-` | Rego ↔ Manifest |

## 규칙 ID 표기 규약

### 단일 SSOT
`<prefix>-<N>` — 접두사는 SSOT 고유(S/D/O/TM/ST/P/F/H/C/M).

### 교차 SSOT
`X<target><source>-<N>`
- `<target>` = LookupKey가 가리키는 SSOT (grounded-against)
- `<source>` = 주장(claim)을 내는 SSOT (claimant)
- SSOT 코드: `O`=OpenAPI, `S`=SSaC, `D`=DDL, `T`=STML, `M`=StateMachine, `P`=Rego, `H`=Hurl, `F`=Func, `N`=Manifest, `A`=Authz, `Q`=sqlc
- 예: SSaC → OpenAPI 방향 → `XOS-`
- 자체검사(same SSOT 내 집계): 동일 문자 반복 (`XOO-`, `XSS-`, `XDD-`, `XPP-`, `XFF-`) — 현재는 per-SSOT 폴더에 흡수됨

## Toulmin 매핑

```
claim   = 검증 대상 (ServiceFunc, PageSpec, sequence, field name, (path, method) 등)
ground  = *rule.Ground (pkg/ground.Build 로 구축, SymbolTable 흡수)
spec    = 규칙별 Spec struct (pkg/rule 공통 또는 고유)

warrant  = 기본 규칙 ("변수가 선언 후 사용되어야 한다")
defeater = 예외 ("currentUser는 암묵적 선언" → IsImplicitVar)
```

## Defeater 전역 카탈로그

| defeater | 면제 warrant | Flags 키 | 면제 조건 |
|----------|-------------|----------|----------|
| `IsPkgModel` | RefExists(SSaC→DDL), CoverageCheck(DDL→SSaC) | `pkgModel` | `pkg/<pkg>/` 제공 내장 모델 테이블 |
| `IsArchived` | CoverageCheck(DDL→SSaC) | `archived` | DDL `@archived` 테이블/컬럼 |
| `IsSensitiveCol` | CoverageCheck(DDL→OpenAPI) | `sensitive` | DDL `-- @sensitive` 컬럼 |
| `IsNoSensitive` | RefExists(sensitive pattern) | `nosensitive` | DDL `-- @nosensitive` 컬럼 |
| `IsImplicitVar` | VarDeclared (S-27~S-30) | `implicit.<name>` | `currentUser`, `request`, `query`, `message` 예약어 |
| `IsSubscribe` | FieldRequired (HTTP 전용), RefExists(request→OpenAPI) | `subscribe` | @subscribe 함수 |
| `IsCustomTS` | RefExists (STML bind→OpenAPI) | `customTS.<name>` | `<page>.custom.ts` 같은 이름 export |

## 검증 흐름

```
ParseAll() → Fullstack (파싱 결과)
  → pkg/ground.Build(fs) → *rule.Ground (Lookup/Types/Pairs/Schemas/Config/Flags)
  → per-SSOT validator 실행
      validate/manifest, /ddl, /openapi, /ssac, /stml, /statemachine, /rego, /func, /hurl
  → per-pair validator 실행
      validate/openapi_ddl, /openapi_ssac, /openapi_hurl, /openapi_manifest,
      /ssac_ddl, /ssac_statemachine, /ssac_func, /ssac_manifest, /ssac_rego, /ssac_authz, /ssac_sqlc,
      /ddl_statemachine, /ddl_rego, /rego_manifest
  → 각 validator 내부:
      rule.Ground + Toulmin Graph (warrant + defeater)
      Graph.Evaluate(claim, ground) per 검증 항목
      verdict + evidence → ValidationError 변환
```

## pkg 구현 상태

| 영역 | internal (확정) | pkg (재구현) | 격차 |
|---|---|---|---|
| ssac/validator | 114파일 + 77 테스트 | 51파일 + 0 테스트 | 규칙 8개 누락 (S-28~30, 43, 51, 53, 55~56), 테스트 0 |
| stml/validator | 59파일 + 17 테스트 | 22파일 + 0 테스트 | 파일 37개 분 로직 미이관, 테스트 0 |
| symbol (심볼 허브) | internal 분산 | `pkg/rule.Ground` + `pkg/ground.Build` | 개념 이관 완료, 완전성 검증 필요 |
| contract 검증 (CT-1~CT-2) | internal에 있음 | **패키지 미존재** | 전체 누락 |
| 교차검증 (구 crosscheck) | internal 수십 테스트 | 14 pair + 5 self (흡수) | 규칙 카탈로그 보존, 테스트 0 |
| 나머지 (manifest/ddl/openapi/…) | — | 이식 완료 | 대부분 적절 |

## internal 일치성 필수 예외

재구현 시 아래 필터/예외는 실사용 피드백으로 확정된 것이므로 반드시 보존 (상세는 각 하위 폴더 README 참조):

| 규칙 | 필수 예외 |
|---|---|
| XSS-11, XDS-12 | `seq.Type == "call"` / `seq.Package != ""` 스킵 |
| XDS-12 | primitive 타입 + sqlc 합성 row type 스킵 |
| XSO-20 | `Result.Wrapper != ""` (Page[T]/Cursor[T]/[]T) 스킵 |
| XDD-61 | `-- @sensitive` / `-- @nosensitive` 보유 컬럼 스킵 |
| XPN-54 | Rego 외 middleware/response 참조도 인정 (검토 중) |
| IsImplicitVar (S-27~30) | `currentUser`/`request`/`query`/`message` 스킵 |
| IsSubscribe | @subscribe 함수는 HTTP 필드 요구 제외 |
| IsCustomTS | `<page>.custom.ts` export 함수 존재 시 bind 통과 |
