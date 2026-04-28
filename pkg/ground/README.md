# pkg/ground

## 변경이력

- 2026-04-28: 4 원칙 준수 형식으로 개정

## 역할

`*yongol.Fullstack` → `*rule.Ground` 어댑터. 파싱 결과에서 검증 규칙이 조회할 데이터를 추출해 공유 `Ground` 에 적재. `pkg/validate` 의 모든 하위 validator (TOULMIN warrant 경로 + plain `g.Lookup[key][name]` 경로) 가 동일 인스턴스를 소비.

> 관련: [`pkg/yongol`](../yongol) (입력 Fullstack), [`pkg/rule`](../rule) (출력 Ground 타입), [`pkg/validate`](../validate) (소비자).

## 공개 함수

| 함수 | 시그니처 | 설명 |
|---|---|---|
| `Build` | `(fs *yongol.Fullstack) *rule.Ground` | 진입점. 빈 Ground 생성 후 populator 순차 호출 (`build.go`) |

## Populator (Build 가 호출)

| 함수 | 소스 | 채우는 버킷 |
|---|---|---|
| `populateOpenAPI` | `OpenAPIDoc` | `Lookup[OpenAPI.{operationId, path, security, method.<path>}]` |
| `populateOpenAPIConstraints` | `RequestConstraints` / `ResponseConstraints` | `Types[OpenAPI.{request,response.constraint}.<op>.{maxLength,format,enum}.<field>]`, `Schemas[...{required,enum,enumFields}]` |
| `populateOpenAPIParams` | `OpenAPIDoc` paths | `Lookup[OpenAPI.param.<op>]`, `Lookup[OpenAPI.request.<op>]`, `Schemas[OpenAPI.response{,.resolved}.<op>]` |
| `populateOpenAPIResponseTypes` | `OpenAPIDoc` 응답 schema | response 타입 정보 |
| `populateSSaC` (+ seq/symbols/request_usage/query_usage) | `ServiceFuncs` | `Lookup[SSaC.{funcName, callRef, modelRef}]`, `Pairs[SSaC.{auth, publish, subscribe}]`, `Schemas[SSaC.response.<funcName>]` |
| `populateVarTypes` | `ServiceFuncs` seqs | `Types[SSaC.var.<funcName>.<var>]` (S-57 등) |
| `populateStates` | `StateDiagrams` | `Lookup[States.{diagram, event.<id>}]` |
| `populateFunc` | `ProjectFuncSpecs` + `YongolPkgSpecs` | `Lookup[Func.spec]`, `Schemas[Func.request.<name>]`, `Types[Func.request.<name>.<field>]` (claims 시 `auth.{issueToken, verifyToken, refreshToken}` 자동 등록) |
| `registerFuncSpec` | (helper) | 단일 func spec 등록 |
| `populateManifest` (+ middleware) | `Manifest` | `Lookup[Manifest.{middleware, claims, claims.keys, roles}]`, `Config[backend.{auth.claims, middleware}]`, `Config[queue.backend]` |
| `populateDDL` (+ archived/check/indexes/varchar) | `DDLTables` | `Lookup[DDL.{table, column.<t>, index.<t>, check.<t>.<col>}]`, `Schemas[DDL.check.<t>.<col>]`, `Types[DDL.varchar.<t>.<col>]` |
| `populateRego` (+ policy) | `ParsedPolicies` | `Pairs[Policy.auth]`, `Lookup[Rego.{claims, roles}]` |
| `populateSymbolTable` / `populateSymbolMethods` | `DDLTables` | `Lookup[SymbolTable.{model, method.<Model>}]` (S-49 — `inflection.Singular` + PascalCase) |
| `populateAuthz` | `Manifest.Authz` | `Lookup[Authz.checkRequest]` (XAS-60) |
| `populateSQLc` | `DDLTables` | `Lookup[SQLc.param.<Model>]` (XQS-14) |
| `populateGoReserved` | (정적) | `Lookup[go.reserved]` |

## 버킷별 용도

| 버킷 | 주 용도 | 소비 |
|---|---|---|
| `Lookup` | 이름 집합 조회 | plain `g.Lookup[key][name]` (RefExists / ForbiddenRef 류) |
| `Pairs` | `"action:resource"` 쌍 매칭 | plain PairMatch (XPS-28 / XSP-29 등) |
| `Types` | 필드·변수 타입 문자열 | plain 비교 (S-57, XDO-69, XFS-44 등) |
| `Schemas` | 정렬된 필드 목록 | plain set diff + `rule.SchemaEvidence` |
| `Config` | 설정 존재 플래그 | `g.Config[key]` |
| `Vars` | SSaC 선언 변수 집합 | populate 안 함 — `pkg/validate/ssac` 가 시퀀스 평가 루프에서 채움. `VarDeclared` (TOULMIN) + `IsImplicitVar` defeater |
| `Flags` | defeater 플래그 | populate 안 함 — rule 평가 컨텍스트에서 `IsSubscribe` / `IsCustomTS` / `IsArchived` 등 설정 |

## 네임스페이스

`Lookup` 의 manifest.yaml 유래 키는 전부 `Manifest.*` 접두사 통일 (`Manifest.middleware`, `Manifest.claims`, `Manifest.claims.keys`, `Manifest.roles`). `g.Config[]` 키 (`backend.auth.claims`, `queue.backend`, `backend.middleware`) 는 "설정 플래그" 의미로 Ground 필드명 유지.

## 의존성

| 패키지 | 역할 |
|---|---|
| `github.com/park-jun-woo/yongol/pkg/yongol` | 입력 Fullstack |
| `github.com/park-jun-woo/yongol/pkg/rule` | 출력 Ground 타입 |
| `github.com/jinzhu/inflection` | `populateSymbolTable` 단수화 |
