# pkg/yongol

## 변경이력

- 2026-04-28: 초기 작성

## 역할

SSOT 오케스트레이션 컨테이너. specs 루트에서 SSOT 디렉토리/파일을 3-상태 (Absent/Declared/Populated) 로 탐지(`DetectSSOTs`) 하고, 탐지된 SSOT 를 1회 파싱(`ParseAll`) 해 `Fullstack` 구조체 하나에 모은다. 이후 모든 검증·체인·코드젠은 이 단일 Fullstack 인스턴스를 공유 소비한다.

## 공개 함수

| 함수 | 시그니처 | 설명 |
|---|---|---|
| `DetectSSOTs` | `DetectSSOTs(root string) ([]DetectedSSOT, error)` | specs 루트를 스캔해 manifest / OpenAPI / DDL / SSaC / States / Policy / Scenario / Func / TSX 의 Presence 를 탐지. SSOTAbsent 는 결과 슬라이스에 포함하지 않음 |
| `ParseAll` | `ParseAll(root string, detected []DetectedSSOT) *Fullstack` | 탐지된 SSOT 를 모두 파싱해 `Fullstack` 반환. 파서 진단은 `fs.ParseDiagnostics` 에 누적 (collect-and-continue) |
| `Fullstack.SetGround` | `(fs *Fullstack) SetGround(g *rule.Ground)` | 외부에서 빌드한 `*rule.Ground` 를 주입 (cycle 회피용) |
| `Fullstack.Ground` | `(fs *Fullstack) Ground() *rule.Ground` | 적재된 `*rule.Ground` 반환 — nil 이면 `SetGround` 선행 필요 |
| `Fullstack.PresenceOf` | `(fs *Fullstack) PresenceOf(kind SSOTKind) SSOTPresence` | SSOTKind → Presence 조회. 미등록 시 `SSOTAbsent` |

## 공개 구조체 / enum

| 타입 | 설명 |
|---|---|
| `Fullstack` | 모든 SSOT 파싱 결과 컨테이너. 주요 필드 — `SpecsDir`, `Manifest`, `OpenAPIDoc`, `OpenAPILines`, `DDLResults` / `DDLTables`, `Policies` / `ParsedPolicies`, `ServiceFuncs`, `StateDiagrams`, `HurlEntries` / `HurlFiles`, `ProjectFuncSpecs`, `YongolPkgSpecs`, `SQLcQueries`, `TSXPages`, `RequestConstraints` / `ResponseConstraints`, `SsacInterfaces`, `Presences`, `ParseDiagnostics` |
| `DetectedSSOT` | 탐지된 SSOT 한 건 — `Kind`, `Path` (절대 경로), `Presence` |
| `SSOTKind` (string) | `KindOpenAPI`, `KindDDL`, `KindSSaC`, `KindStates`, `KindPolicy`, `KindScenario`, `KindFunc`, `KindConfig`, `KindTSX` |
| `SSOTPresence` (int) | `SSOTAbsent` (0), `SSOTDeclared` (1), `SSOTPopulated` (2) |
