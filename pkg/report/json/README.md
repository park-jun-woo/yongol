# pkg/report/json

## 변경이력

- 2026-04-28: 초기 작성

## 역할

`validate.Report` 를 yongol bespoke flat JSON 으로 직렬화. `yongol validate --format json` 의 출력 포맷이며, 진단 메시지 앞 `[RULE-ID]` 토큰을 `rule_id` 필드로 분리하고 ERROR/WARNING/checks 카운트를 `summary` 에 집계한다.

> 상위 문서: [`pkg/report/README.md`](../README.md)

## 공개 함수 / 구조체

| 식별자 | 종류 | 설명 |
|---|---|---|
| `Emit(report, yongolVersion, specsDir, checks)` | func | `validate.Report` → flat JSON `[]byte` 변환 |
| `Document` | type | top-level 문서 (`yongol_version`, `specs_dir`, `summary`, `diagnostics`) |
| `Summary` | type | `errors` / `warnings` / `checks` 카운트 |
| `Diagnostic` | type | 개별 진단 (`rule_id` / `level` / `file` / `line` / `col` / `message`) |

## 보조 헬퍼

`extractRuleID`, `appendDiagnostic`, `collectReportDiagnostics`, `relativeFile`, `tryAbsRelativeFile`.
