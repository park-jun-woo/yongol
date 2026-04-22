# pkg/parser

> 모든 공개 파서는 **`(result, []diagnostic.Diagnostic)`** 시그니처를 따른다. 에러는 `error` 대신 `pkg/diagnostic` 의 Diagnostic 슬라이스로 누적 보고한다.

## 자체 파서

| 패키지 | 입력 | 출력 | entry | 설명 |
|--------|------|------|-------|------|
| `ssac/` | `.ssac` | `[]ServiceFunc` | `ParseDir(dir)`, `ParseFile(path)` | 서비스 함수 시퀀스 파싱 (Go AST 활용) |
| `stml/` | `.html` | `[]PageSpec` | `ParseDir(dir)`, `ParseFile(path)`, `ParseReader(name, r)` | 프론트엔드 페이지 템플릿 파싱 (x/net/html 활용) |
| `statemachine/` | Mermaid `.md` | `[]*StateDiagram` | `ParseDir(dir)`, `ParseFile(path)`, `Parse(id, content, file)` | 상태 전이 다이어그램 파싱 |
| `funcspec/` | Go `.go` | `[]FuncSpec` | `ParseDir(dir)`, `ParseFile(path)` | 커스텀 함수 스펙 파싱 (Go AST 활용) |
| `hurl/` | `.hurl` | `[]HurlEntry` | `ParseFile(path)` | 통합 테스트 시나리오 파싱. 파일 단위만 제공, 디렉토리 순회는 호출측 책임 |
| `manifest/` | `manifest.yaml` | `*ProjectConfig` | `Load(specsDir)` | 프로젝트 설정 파싱 (yaml.v3 활용) |
| `openapi/` | `*openapi3.T` | `map[op][field]FieldConstraint` | `ExtractRequestConstraints(doc)`, `ExtractResponseConstraints(doc)` | OpenAPI 요청/응답 필드의 type/format/length/enum/required 제약 추출. 에러 미발생 가정 (doc 이미 파싱됨) |
| `toulmin/` | Go `.go` | `*Graph` | (예정) | **🚧 작업중** — Toulmin 규칙 그래프 파서. 상용 엔진 정립 전까지 보류. 현재 규칙 엔진은 OPA Rego 만 지원 |

## 구조화 파서 + 외부 검증

| 패키지 | 구조화 파서 | 출력 | 외부 검증 | 설명 |
|--------|-----------|------|----------|------|
| `ddl/` | `ParseTables(dir)` | `[]Table` | `ParseDir(dir)` → `[]*pg_query.ParseResult` | DDL 테이블/컬럼/FK/인덱스/CHECK |
| `rego/` | `ParsePolicies(dir)` | `[]Policy` | `ParseDir(dir)` → `[]*ast.Module`, `ParsePolicyFile(path)` | allow 규칙, @ownership, claims 참조 |

## 외부 라이브러리 (래퍼 없이 직접 사용)

| 라이브러리 | 대상 | 출력 | 설명 |
|-----------|------|------|------|
| `kin-openapi` | OpenAPI YAML | `*openapi3.T` | API 엔드포인트 스키마. `pkg/yongol.ParseAll` 에서 `openapi3.NewLoader().LoadFromFile()` 로 로드 |

## 관련 상위 패키지

| 패키지 | 역할 |
|--------|------|
| `pkg/yongol` | 모든 SSOT 파싱 결과를 담는 `Fullstack` 컨테이너 + `ParseAll(root, detected, skip)`. `pkg/parser/*` 를 조합 호출 |
| `pkg/diagnostic` | 파서/검증기 공용 진단 타입 (`Diagnostic`, `Phase`, `Level`, `Loc`) |

## DDL 구조체

```go
type Table struct {
    Name        string
    Columns     map[string]string   // column → Go type
    ColumnOrder []string            // DDL 정의 순서
    ForeignKeys []ForeignKey        // {Column, RefTable, RefColumn}
    Indexes     []Index             // {Name, Columns, IsUnique}
    PrimaryKey  []string
    VarcharLen  map[string]int      // column → VARCHAR(N)
    CheckEnums  map[string][]string // column → CHECK IN values
}
```

## Rego 구조체

```go
type Policy struct {
    File       string
    Rules      []AllowRule          // {Actions, Resource, UsesOwner, UsesRole, RoleValue}
    Ownerships []OwnershipMapping   // {Resource, Table, Column, JoinTable, JoinFK}
    ClaimsRefs []string             // input.claims.xxx 참조 (중복 제거)
}
```

## OpenAPI 구조체

```go
type FieldConstraint struct {
    Type      string
    Format    string
    MaxLength *int     // nil = 무제한
    MinLength *int
    Enum      []string
    Required  bool
}
```

## Hurl 구조체

```go
type HurlEntry struct {
    Method     string
    Path       string
    StatusCode string
    // + 추가 필드
}
```
