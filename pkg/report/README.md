# pkg/report

## 변경이력

- 2026-04-28: 초기 작성

## 역할

`yongol validate` / `yongol generate` 의 실행 결과 (`validate.Report`) 를 외부 포맷으로 직렬화하는 emitter 모음. 본 폴더에는 Go 파일이 없고, 자식 패키지가 포맷별 구현을 담당한다.

## 자식 패키지

| 패키지 | 포맷 | 진입점 |
|---|---|---|
| [`json/`](json/) | yongol bespoke flat JSON | `json.Emit(report, yongolVersion, specsDir, checks) ([]byte, error)` |
| [`sarif/`](sarif/) | SARIF 2.1.0 (GitHub Code Scanning 호환) | `sarif.Emit(report, yongolVersion, specsDir, *catalog.Catalog) ([]byte, error)` |

두 emitter 모두 진단 메시지 앞 `[RULE-ID]` 토큰을 추출해 `rule_id` / `ruleId` 필드로 분리하고, `specsDir` 기준 상대 경로 + slash 형식으로 파일 위치를 정규화한다.
