# pkg/report/sarif

## 변경이력

- 2026-04-28: 초기 작성

## 역할

`validate.Report` 를 SARIF 2.1.0 문서로 직렬화. GitHub Code Scanning / 다른 SARIF consumer 가 yongol 진단을 그대로 인입할 수 있도록 한다. `catalog` 가 주어지면 rulebook 전체 규칙을 `tool.driver.rules[]` 에 방출하고, 각 result 의 `ruleIndex` 를 카탈로그 인덱스와 역참조 가능하게 연결한다.

> 상위 문서: [`pkg/report/README.md`](../README.md)

## 공개 함수 / 구조체

| 식별자 | 종류 | 설명 |
|---|---|---|
| `Emit(report, yongolVersion, specsDir, *catalog.Catalog)` | func | `validate.Report` + 카탈로그 → SARIF 2.1.0 JSON `[]byte` |
| `Document` | type | SARIF top-level (`$schema`, `version`, `runs`) |
| `Run` | type | 단일 도구 실행 (`tool` + `results`) |
| `Tool` / `Driver` | type | 도구 식별자 (name / version / informationUri / rules) |
| `Rule` | type | 규칙 메타 (id / name / shortDescription / helpUri / defaultConfiguration) |
| `DefaultConfiguration` | type | 규칙 기본 severity (rulebook Level → SARIF level) |
| `Result` | type | 단일 진단 (ruleId / ruleIndex / level / message / locations) |
| `Location` / `PhysicalLocation` / `ArtifactLocation` / `Region` | type | 위치 표현 (1-based line/column) |
| `Message` | type | 사람이 읽는 진단 텍스트 래퍼 |

## 보조 헬퍼

`collectResults`, `buildResult`, `buildResultLocations`, `buildDriverRules`, `ruleFromMeta`, `attachRuleIndex`, `mapLevel`, `extractRuleID`, `relativeArtifactURI`, `tryAbsRelativeURI`, `regionOrNil`, `sortStrings`.
