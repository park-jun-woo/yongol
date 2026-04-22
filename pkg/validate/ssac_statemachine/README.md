# pkg/validate/ssac_statemachine

SSaC @state 선언과 Mermaid stateDiagram 간의 diagram/event/guard 일관성 검증.

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = defeater 실 작동 또는 반례 확장 가능 / `IF-ELSE` = 단일 판정·Ground 조회 — 본 폴더는 전부 IF-ELSE

## RefExists (IF-ELSE)

| 규칙 ID | LookupKey | 설명 | 구현 방식 |
|---------|-----------|------|----------|
| XMS-24 | `States.diagram` | SSaC @state → diagram 존재 | IF-ELSE |
| XMS-25 | `States.event.<diagramID>` | @state transition → diagram event | IF-ELSE |
| XSM-23 | `SSaC.funcName` | States transition event → SSaC | IF-ELSE |

## 고유 함수 (IF-ELSE)

| 규칙 ID | 함수명 | 설명 | 구현 방식 |
|---------|--------|------|----------|
| XSM-26 | `MissingStateGuard` | 상태 전이 참여하는데 @state 없음 (WARNING) | IF-ELSE |

## Defeater

없음.
