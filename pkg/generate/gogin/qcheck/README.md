# pkg/generate/gogin/qcheck

## 변경이력

- 2026-04-28: 초기 작성

## 역할

생성된 Go 소스에 대해 filefunc 품질 규칙 (Q1 nesting depth, Q4 range body 순수 라인 수) 위반과 방어 패턴 누락 (DF-01 unmarshal 무체크 / DF-02 Scan 무체크 / DF-06 defer Close 누락) 을 AST 로 스캔해 WARN 메시지를 반환한다.

## 진입점

| 함수 | 시그니처 | 설명 |
|---|---|---|
| `WarnExceeds` | `(filename, src string, lim Limits) []string` | Q1/Q4 위반 + DF-01/02/06 finding 을 WARN 라인 리스트로 반환 |
| `MeasureDepth` | `(filename, src string) ([]DepthReport, error)` | 함수별 최대 nesting depth (Q1 근거) |
| `MeasurePureLines` | `(filename, src string) ([]PureLinesReport, error)` | 루프 body 순수 라인 수 (Q4 근거) |
| `ScanUncheckedUnmarshal` | `(filename, src string) ([]DefensiveFinding, error)` | DF-01 |
| `ScanUncheckedScan` | `(filename, src string) ([]DefensiveFinding, error)` | DF-02 |
| `ScanMissingDeferClose` | `(filename, src string) ([]DefensiveFinding, error)` | DF-06 |
| `DefaultLimits` | `() Limits` | filefunc 기본 budget (depth ≤ 3, pure ≤ 10) |

## 공개 구조체

| 타입 | 설명 |
|---|---|
| `Limits` | Q1 depth + Q4 순수 라인 상한 |
| `DepthReport` | 한 함수의 최대 nesting depth |
| `PureLinesReport` | 한 루프 body 의 순수 라인 수 |
| `DefensiveFinding` | 파일:라인 + 카테고리 (DF-01/02/06) |
