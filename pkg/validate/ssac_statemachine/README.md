# pkg/validate/ssac_statemachine

## 변경이력

- 2026-04-28: 4 원칙 준수 형식으로 개정

## 역할

SSaC `@state` 선언과 Mermaid stateDiagram 간 diagram/event/guard 일관성 검증 (XMS-*, XSM-*).

> 상위 문서: [`pkg/validate/README.md`](../README.md)
> **구현 방식 범례**: `TOULMIN` = defeater 작동 / `IF-ELSE` = 단일 판정·Ground 조회 — 본 폴더는 전부 IF-ELSE

## 검증 규칙

| 규칙 ID | 함수명 | 설명 | 구현 방식 | pkg 구현 |
|---|---|---|---|---|
| XMS-24 | `StateDiagramExists` | `@state` DiagramID → diagram 존재 (ERROR) | IF-ELSE | ✓ |
| XMS-25 | `StateEvent` | `@state` transition → diagram event 정의 (ERROR) | IF-ELSE | ✓ |
| XSM-23 | `TransitionToFunc` | diagram transition event → SSaC 함수 (ERROR) | IF-ELSE | ✓ |
| XSM-26 | `MissingStateGuard` | 전이 참여 함수에 `@state` 없음 (WARNING) | IF-ELSE | ✓ |
| XSM-27 | `StateIntentDeclaration` | stateful POST/PUT/DELETE 는 `@state` 또는 `// @state-neutral` 의도 선언 강제 (WARNING) | IF-ELSE | ✓ |

## Defeater

없음. XSM-27 은 stateful 리소스 (path segment ↔ diagram 매핑) 가 식별될 때만 발화 — `is_stateful_resource.go`.
