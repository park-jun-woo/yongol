# pkg/ground

`*yongol.Fullstack` → `*rule.Ground` 어댑터. 파싱 결과에서 검증 규칙이 조회할 데이터를 추출해 공유 `Ground`에 적재한다.

`pkg/validate`의 모든 하위 validator(TOULMIN 35개 + plain lookup 107개)가 동일한 `Ground` 인스턴스를 소비한다. Toulmin warrant 경로든 plain `g.Lookup[key][name]` 경로든 같은 버킷을 참조.

## API

```go
func Build(fs *yongol.Fullstack) *rule.Ground
```

빈 `Ground`를 만들고 populator를 순차 호출해 반환. 진입점은 `build.go` 참조.

## Populator

| 함수 | 소스 | 채우는 Ground 버킷 |
|---|---|---|
| `populateOpenAPI` | `OpenAPIDoc` | `Lookup[OpenAPI.operationId]`, `Lookup[OpenAPI.path]`, `Lookup[OpenAPI.security]`, `Lookup[OpenAPI.method.<path>]` |
| `populateSSaC` | `ServiceFuncs` | `Lookup[SSaC.funcName]`, `Lookup[SSaC.callRef]`, `Lookup[SSaC.modelRef]`, `Pairs[SSaC.auth]`, `Pairs[SSaC.publish]`, `Pairs[SSaC.subscribe]`, `Schemas[SSaC.response.<funcName>]` |
| `populateStates` | `StateDiagrams` | `Lookup[States.diagram]`, `Lookup[States.event.<id>]` |
| `populateFunc` | `ProjectFuncSpecs`+`YongolPkgSpecs` | `Lookup[Func.spec]`, `Schemas[Func.request.<name>]`, `Types[Func.request.<name>.<field>]` (+ claims 설정 시 `auth.issueToken`/`verifyToken`/`refreshToken` 자동 등록) |
| `populateManifest` | `Manifest` | `Lookup[Manifest.middleware]`, `Lookup[Manifest.claims]`, `Lookup[Manifest.claims.keys]`, `Lookup[Manifest.roles]`, `Config[backend.auth.claims]`, `Config[queue.backend]`, `Config[backend.middleware]` |
| `populateDDL` | `DDLTables` | `Lookup[DDL.table]`, `Lookup[DDL.column.<t>]`, `Lookup[DDL.index.<t>]`, `Lookup[DDL.check.<t>.<col>]`, `Schemas[DDL.check.<t>.<col>]`, `Types[DDL.varchar.<t>.<col>]` |
| `populateRego` | `ParsedPolicies` | `Pairs[Policy.auth]`, `Lookup[Rego.claims]`, `Lookup[Rego.roles]` |
| `populateOpenAPIConstraints` | `RequestConstraints`/`ResponseConstraints` | `Types[OpenAPI.{request,response.constraint}.<op>.{maxLength,format,enum}.<field>]`, `Schemas[OpenAPI.{request,response.constraint}.<op>.{required,enum,enumFields}]` |
| `populateOpenAPIParams` | `OpenAPIDoc` paths | `Lookup[OpenAPI.param.<op>]`, `Lookup[OpenAPI.request.<op>]`, `Schemas[OpenAPI.response{,.resolved}.<op>]` |
| `populateSymbolTable` | `DDLTables` | `Lookup[SymbolTable.model]` (테이블명 → `inflection.Singular` + PascalCase) |
| `populateSymbolMethods` | `DDLTables` | `Lookup[SymbolTable.method.<Model>]` (표준 CRUD 메서드 set — S-49) |
| `populateAuthz` | `Manifest.Authz` | `Lookup[Authz.checkRequest]` (built-in authz.CheckRequest 필드 set — XAS-60) |
| `populateSQLc` | `DDLTables` | `Lookup[SQLc.param.<Model>]` (컬럼명 set — XQS-14) |
| `populateSSaCRequestUsage` | `ServiceFuncs` | `Lookup[SSaC.requestUsage.<funcName>]` (함수별 request 필드 사용 — S-51) |
| `populateSSaCQueryUsage` | `ServiceFuncs` | `Lookup[SSaC.queryUsage]` (전역 query 필드 사용 — S-53) |
| `populateVarTypes` | `ServiceFuncs` seqs | `Types[SSaC.var.<funcName>.<var>]` (SSaC @call arg type 검증에 사용) |
| `populateGoReservedWords` | (정적) | `Lookup[go.reserved]` |

## 버킷별 용도

| 버킷 | 주 용도 | 소비 경로 |
|---|---|---|
| `Lookup` | 이름 집합 조회 | 대부분 plain `g.Lookup[key][name]` if-else (RefExists/ForbiddenRef 류) |
| `Pairs` | `"action:resource"` 쌍 매칭 | plain PairMatch (XPS-28/XSP-29 등) |
| `Types` | 필드·변수 타입 문자열 | plain 비교 (S-57, XDO-69, XFS-44 등) |
| `Schemas` | 정렬된 필드 목록 | plain set diff + `rule.SchemaEvidence` 반환 |
| `Config` | 설정 존재 플래그 | plain `g.Config[key]` 체크 |
| `Vars` | SSaC 선언 변수 집합 | **populate 안 함** — `pkg/validate/ssac` 가 시퀀스 평가 루프에서 채움. `VarDeclared` (TOULMIN)가 `IsImplicitVar` defeater와 함께 소비 |
| `Flags` | defeater 플래그 | **populate 안 함** — rule 평가 컨텍스트에서 `IsSubscribe`/`IsCustomTS`/`IsArchived` 등 설정 |

> **참고**: `Lookup`/`Pairs`/`Types`/`Schemas`/`Config` 버킷은 TOULMIN 경로(`pkg/rule.RefExists` 등)와 plain lookup 경로 모두에서 동일하게 소비된다. populator는 소비 경로와 무관하게 Ground만 채우면 됨.

## 네임스페이스

`Lookup` 의 manifest.yaml 유래 키는 전부 `Manifest.*` 접두사로 통일 (`Manifest.middleware`, `Manifest.claims`, `Manifest.claims.keys`, `Manifest.roles`). `g.Config[]` 버킷 자체(`backend.auth.claims`, `queue.backend`, `backend.middleware` 키)는 "설정 플래그" 의미로 Ground 필드명 유지.

## 의존성

- `github.com/park-jun-woo/yongol/pkg/yongol` (입력)
- `github.com/park-jun-woo/yongol/pkg/rule` (출력 타입)
- `github.com/jinzhu/inflection` (`populateSymbolTable` 단수화)
