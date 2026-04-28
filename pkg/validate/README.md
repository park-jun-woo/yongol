# pkg/validate

## 변경이력

- 2026-04-28: 4 원칙 준수 형식으로 개정

## 역할

yongol 의 단일·교차 SSOT 정합성 검증 통합 패키지. 도메인별 서브패키지가 `pkg/rule` 공통 함수와 자체 함수로 진단을 emit. 규칙 카탈로그 전체는 [`rulebook.md`](../../rulebook.md).

> **구현 방식 범례**: `TOULMIN` = pkg/rule + defeater / `IF-ELSE` = 단일 흐름 검사
> **Toulmin 매핑**: `claim` = 검증 대상, `ground` = `*rule.Ground` (`pkg/ground.Build`), `warrant` = 기본 규칙, `defeater` = 예외.

## 단일 SSOT 서브패키지

| 패키지 | 규칙 ID 접두사 |
|---|---|
| [`manifest/`](manifest/) | `C-`, `CORS-`, `OBS-`, `SEC-2/3/4`, `XN*-90` |
| [`ddl/`](ddl/) | `D-`, `XDD-` |
| [`query/`](query/) | `Q-` |
| [`openapi/`](openapi/) | `O-`, `XOO-` |
| [`ssac/`](ssac/) | `S-`, `XSS-` |
| [`statemachine/`](statemachine/) | `ST-` |
| [`rego/`](rego/) | `P-`, `XPP-` |
| [`funcspec/`](funcspec/) | `F-`, `XFF-` |
| [`hurl/`](hurl/) | `H-` |
| [`tsx/`](tsx/) | `TX-` |
| [`migration/`](migration/) | `MG-` |
| [`contract/`](contract/) | `CT-`, `PRV-` |

## 교차 SSOT 서브패키지 (X-prefix 룰)

| 폴더 | 규칙 ID 접두사 |
|---|---|
| [`openapi_ddl/`](openapi_ddl/) | `XDO-`, `XOD-` |
| [`openapi_ssac/`](openapi_ssac/) | `XOS-`, `XSO-` |
| [`openapi_manifest/`](openapi_manifest/) | `XNO-`, `XON-`, `SEC-04` |
| [`manifest_ddl/`](manifest_ddl/) | `XDN-` |
| [`ssac_ddl/`](ssac_ddl/) | `XSD-`, `XDS-` |
| [`ssac_sqlc/`](ssac_sqlc/) | `XQS-` |
| [`ssac_func/`](ssac_func/) | `XFS-`, `XSF-` |
| [`ssac_manifest/`](ssac_manifest/) | `XNS-` |
| [`ssac_rego/`](ssac_rego/) | `XPS-`, `XSP-` |
| [`ssac_statemachine/`](ssac_statemachine/) | `XMS-`, `XSM-` |
| [`ssac_authz/`](ssac_authz/) | `XAS-` |
| [`ddl_rego/`](ddl_rego/) | `XDP-` |
| [`ddl_statemachine/`](ddl_statemachine/) | `XDM-` |
| [`rego_manifest/`](rego_manifest/) | `XNP-`, `XPN-` |
| [`query_rego/`](query_rego/) | `XQP-` |
| [`hurl_openapi/`](hurl_openapi/) | `XOH-` |
| [`hurl_statemachine/`](hurl_statemachine/) | `XMH-` |
| [`hurl_manifest/`](hurl_manifest/) | `XNH-` |
| [`tsx_openapi/`](tsx_openapi/) | `XOT-` |

## 규칙 ID 표기 규약

- **단일 SSOT**: `<prefix>-<N>` — 접두사 SSOT 고유 (S/D/O/TM/ST/P/F/H/C/M).
- **교차 SSOT**: `X<target><source>-<N>` — `<target>` = LookupKey 의 SSOT (grounded-against), `<source>` = 주장 SSOT (claimant).
- SSOT 코드: `O`=OpenAPI, `S`=SSaC, `D`=DDL, `T`=STML/TSX, `M`=StateMachine, `P`=Rego, `H`=Hurl, `F`=Func, `N`=Manifest, `A`=Authz, `Q`=sqlc.

## Defeater 전역 카탈로그

| defeater | 면제 warrant | 면제 조건 |
|---|---|---|
| `IsPkgModel` | RefExists/CoverageCheck (SSaC↔DDL) | `pkg/<pkg>/` 제공 내장 모델 테이블 |
| `IsArchived` | CoverageCheck(DDL→SSaC) | DDL `@archived` 테이블/컬럼 |
| `IsSensitiveCol` | CoverageCheck(DDL→OpenAPI) | DDL `-- @sensitive` 컬럼 |
| `IsNoSensitive` | RefExists(sensitive pattern) | DDL `-- @nosensitive` 컬럼 |
| `IsImplicitVar` | VarDeclared (S-27~30) | `currentUser`/`request`/`query`/`message` 예약어 |
| `IsSubscribe` | FieldRequired (HTTP 전용) | `@subscribe` 함수 |
| `IsCustomTS` | RefExists (STML bind→OpenAPI) | `<page>.custom.ts` export |
| `IsNullableIntentional` | XDO-76 | DDL `-- @nullable` 어노테이션 |

## 검증 흐름

```
ParseAll() → Fullstack
  → pkg/ground.Build(fs) → *rule.Ground
  → per-SSOT validator (manifest/ddl/query/openapi/ssac/statemachine/rego/funcspec/hurl/tsx/migration/contract)
  → per-pair validator (openapi_*, ssac_*, ddl_*, rego_*, query_*, hurl_*, tsx_*)
  → 각 validator: rule.Ground + Toulmin Graph (warrant + defeater) → diagnostic
```

## internal 일치성 필수 예외

각 하위 폴더 README 의 *internal 일치성 메모* 절 참조. 핵심:

| 규칙 | 필수 예외 |
|---|---|
| XSS-11, XDS-12 | `seq.Type == "call"` / `seq.Package != ""` 스킵 |
| XDS-12 | primitive 타입 + sqlc 합성 row type 스킵 |
| XDD-61 | `-- @sensitive` / `-- @nosensitive` 컬럼 스킵 |
| XPN-54 | Rego 외 middleware/response 참조도 인정 |
| IsImplicitVar (S-27~30) | `currentUser`/`request`/`query`/`message` 스킵 |
| IsSubscribe | `@subscribe` 함수는 HTTP 필드 요구 제외 |
| IsCustomTS | `<page>.custom.ts` export 함수 존재 시 bind 통과 |
