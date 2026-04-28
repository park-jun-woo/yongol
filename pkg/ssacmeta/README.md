# pkg/ssacmeta

## 변경이력

- 2026-04-28: 초기 작성

## 역할

ssac 형제 repo (`github.com/park-jun-woo/ssac`) 의 `pkg/<name>/interface.yaml` 메타데이터를 yongol 로 로드한다. 각 ssac 패키지(cache / session / queue / auth / authz …)가 선언한 DB-access port / sqlc 쿼리 스펙 / canonical DDL·queries 경로를 yongol 의 검증·코드젠 (Phase002/004/005) 이 소비한다.

## 공개 함수

| 함수 | 시그니처 | 설명 |
|---|---|---|
| `LoadPackageInterface` | `LoadPackageInterface(path string) (*PackageInterface, error)` | 단일 `interface.yaml` 파일을 파싱. 파일 부재 시 `(nil, nil)` |
| `LoadPackageInterfaces` | `LoadPackageInterfaces(ssacRoot string) (map[string]*PackageInterface, error)` | `<ssacRoot>/pkg/*/interface.yaml` 을 모두 로드해 `package:` 키 맵 반환 |
| `EvaluateWhen` | `EvaluateWhen(expr string, manifest map[string]any) bool` | `interface.yaml` 의 `when:` 표현식 평가 — `always`, `manifest.<path> == "<value>"`, `manifest.<path>` (truthy) 지원 |

## 공개 구조체

| 타입 | 설명 |
|---|---|
| `PackageInterface` | `interface.yaml` 루트. `Version`, `Package`, `Description`, `Ports`, `PortsMeta`, `CanonicalDDL`, `CanonicalQueries`, `SourcePath` |
| `Port` | 단일 DB-access port. `Name`, `Description`, `When`, `UsedBy`, `Query` |
| `QuerySpec` | sqlc 쿼리 스펙. `Cardinality` (one/many/exec), `Params`, `Returns` |
| `Field` | params / returns 의 단일 필드 — `Name`, `Type` |
| `Meta` | dynamic-port 패키지 (authz 등) 의 `ports_meta` — `Rule`, `Description`, `Convention` |
| `Convention` | dynamic-port 네이밍·shape 규약 — `Name`, `Cardinality`, `Params`, `Returns` |
