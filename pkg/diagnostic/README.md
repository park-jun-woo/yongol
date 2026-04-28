# pkg/diagnostic

## 변경이력

- 2026-04-28: 초기 작성

## 역할

파서·검증기·교차검증·코드젠 단계에서 공통으로 사용하는 진단 메시지 타입. `error` 대신 `[]Diagnostic` 슬라이스로 누적해 보고하는 collect-and-continue 모델의 데이터 형식.

## 공개 구조체

| 타입 | 설명 |
|---|---|
| `Diagnostic` | 단일 진단. `File`, `Line`, `Phase`, `Level`, `Message` (Rule-ID + 무엇이 잘못되었는지), `Advice` (어떻게 고치는지 — 빈 문자열이면 출력 시 숨김) |

## 공개 enum 타입

| 타입 | 상수 | 설명 |
|---|---|---|
| `Phase` (string alias) | `PhaseParse` (`"parse"`), `PhaseValidate` (`"validate"`) | 진단을 발생시킨 단계 |
| `Level` (string alias) | `LevelError` (`"ERROR"`), `LevelWarning` (`"WARNING"`) | 심각도 |

## 비고

`Phase` / `Level` 모두 `type X string` alias 라 `string(PhaseParse)` 왕복 변환이 가능하다. 리포트 포맷팅은 `pkg/report` 에서 처리한다.
