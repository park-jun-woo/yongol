# pkg/parser

## 변경이력

- 2026-04-28: 4 원칙 준수 형식으로 개정

## 역할

9 SSOT 의 텍스트 입력을 구조체로 변환하는 파서 모음. 모든 공개 파서는 `(result, []diagnostic.Diagnostic)` 시그니처를 따르며 에러는 `pkg/diagnostic` 슬라이스로 누적 보고.

> 상위: [`pkg/yongol/README.md`](../yongol/README.md) (`Fullstack` 컨테이너 + `ParseAll` 호출자)

## 자체 파서 (sub-package)

| 패키지 | 입력 → 출력 | entry |
|---|---|---|
| `ssac/` | `.ssac` → `[]ServiceFunc` | `ParseDir(dir)`, `ParseFile(path)` |
| `stml/` | `.html` → `[]PageSpec` | `ParseDir(dir)`, `ParseFile(path)`, `ParseReader(name, r)` |
| `statemachine/` | Mermaid `.md` → `[]*StateDiagram` | `ParseDir(dir)`, `ParseFile(path)`, `Parse(id, content, file)` |
| `funcspec/` | Go `.go` → `[]FuncSpec` | `ParseDir(dir)`, `ParseFile(path)` |
| `hurl/` | `.hurl` → `[]HurlEntry` | `ParseFile(path)` (디렉토리 순회는 호출측 책임) |
| `manifest/` | `manifest.yaml` → `*ProjectConfig` | `Load(specsDir)` (yaml.v3) |
| `openapi/` | `*openapi3.T` → `map[op][field]FieldConstraint` | `ExtractRequestConstraints(doc)`, `ExtractResponseConstraints(doc)` |
| `ddl/` | DDL `.sql` → `[]Table` + `[]*pg_query.ParseResult` | `ParseTables(dir)`, `ParseDir(dir)` |
| `rego/` | OPA `.rego` → `[]Policy` + `[]*ast.Module` | `ParsePolicies(dir)`, `ParseDir(dir)`, `ParsePolicyFile(path)` |
| `sqlc/` | `sqlc.yaml` → 메타 | sqlc generate 보조 (`pkg/parser/sqlc/`) |

## 공개 구조체 (요약)

| 타입 | 위치 | 설명 |
|---|---|---|
| `Table` | `ddl/` | name, columns, FK, indexes, PK, varcharLen, checkEnums |
| `Policy` | `rego/` | rules (`AllowRule`), ownerships (`OwnershipMapping`), claimsRefs |
| `FieldConstraint` | `openapi/` | type, format, maxLength, minLength, enum, required |
| `HurlEntry` | `hurl/` | method, path, statusCode + 추가 필드 |
| `StateDiagram` / `Transition` | `statemachine/` | mermaid stateDiagram 구조 |
| `ServiceFunc` | `ssac/` | SSaC 함수 시퀀스 |
| `PageSpec` | `stml/` | STML 페이지 메타 |
| `FuncSpec` | `funcspec/` | Go AST 함수 스펙 |
| `ProjectConfig` | `manifest/` | manifest.yaml 루트 |

## 외부 라이브러리 (직접 사용)

| 라이브러리 | 용도 |
|---|---|
| `kin-openapi` (`openapi3.T`) | OpenAPI YAML 로드 — `pkg/yongol.ParseAll` 에서 `openapi3.NewLoader().LoadFromFile()` |
| `pganalyze/pg_query_go` | DDL AST |
| `open-policy-agent/opa/ast` | Rego AST |
| `gopkg.in/yaml.v3` | manifest |
| `golang.org/x/net/html` | stml |
