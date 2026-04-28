# pkg/validate/hurl_statemachine

## 변경이력

- 2026-04-28: 초기 작성

## 역할

Hurl 파일 내 operation 호출 순서가 Mermaid stateDiagram 의 전이 규칙을 위반하지 않는지 교차 검증한다. 도달 가능한 상태 집합을 entry 단위로 갱신하며, 현재 상태에서 허용되지 않는 전이가 발생하면 WARNING 을 emit.

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `IF-ELSE` = state machine reachable-set 시뮬레이션

## 검증 규칙

| 규칙 ID | 함수명 | 설명 | 구현 방식 | pkg 구현 |
|---|---|---|---|---|
| XOH-05 | `xoh_05_state_transition_order` | 같은 파일 내 operation 호출 순서가 state machine 전이 규칙을 위반하지 않음 (WARNING) | IF-ELSE | ✓ |

## 주요 함수

| 함수 | 설명 |
|---|---|
| `Run(fs)` | Hurl ↔ StateMachine 교차 검증 실행 (XOH-05) |

## 보조 헬퍼

`buildOpIDLookup` (METHOD /path → operationId), `groupByFile`, `checkFileOrder`, `checkDiagramOrder`, `inspectDiagramEntry`, `advanceReachable`, `transitionFrom`, `normPath`, `normPathSegment`.
