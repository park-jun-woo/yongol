# pkg/parser/rego

## 변경이력

- 2026-04-28: 초기 작성

## 역할

OPA Rego 정책 파일 (`*.rego`) 디렉토리를 파싱한다. OPA AST 기반 strict parse 로 R-1 문법 ERROR 를 잡고, 정규식 기반 메타 추출로 `allow` 규칙 / `@ownership` 어노테이션 / `input.claims.*` 참조를 수집한다. AST 모듈 결과와 구조화된 `Policy` 결과 두 가지를 함께 노출한다.

## 진입점

| 함수 | 시그니처 | 설명 |
|---|---|---|
| `ParsePolicies` | `(dir string) ([]Policy, []diagnostic.Diagnostic)` | 디렉토리 내 `*.rego` 에서 구조화된 Policy 추출 |
| `ParsePolicyFile` | `(path string) (*Policy, []diagnostic.Diagnostic)` | 단일 `.rego` 파일 파싱 (OPA strict parse + 메타 추출) |
| `ParseDir` | `(dir string) ([]*ast.Module, []diagnostic.Diagnostic)` | 디렉토리 내 `*.rego` 를 `opa/ast` 로 파싱 |

## 공개 구조체

| 타입 | 설명 |
|---|---|
| `Policy` | 정책 파싱 결과 (`File / Rules / Ownerships / ClaimsRefs`) |
| `AllowRule` | allow 규칙에서 추출한 액션-리소스 쌍 (`Actions / Resource / UsesOwner / UsesRole / RoleValue / SourceLine`) |
| `OwnershipMapping` | `@ownership` 어노테이션 (`Resource / Table / Column / JoinTable / JoinFK / SourceLine`) |

## 비고

- OPA AST 라이브러리 (`github.com/open-policy-agent/opa/ast`) 직접 사용.
- `@ownership` 형식: `<resource>: <table>.<column>` 또는 `<resource>: <table>.<column> via <join_table>.<join_fk>`.
