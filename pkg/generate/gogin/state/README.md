# pkg/generate/gogin/state

## 변경이력

- 2026-04-28: 4 원칙 준수 형식으로 개정

## 역할

Mermaid stateDiagram → 상태 전이 맵 + `<ID>CanTransition` / `<ID>NextState` 런타임 guard 함수 생성. SSaC `@state` 시퀀스가 호출. DB 접근/외부 import 없는 순수 in-memory map.

> 상위: [`pkg/generate/gogin/README.md`](../README.md) [7]. 활성 조건: `len(fs.StateDiagrams) > 0`.

## 공개 함수

| 함수 | 시그니처 | 설명 |
|---|---|---|
| `Generate` | `(fs *yongol.Fullstack, artifactsDir string) error` | 진입점. diagram 별 1 파일 emit |
| `buildTransitionMap` | `(diagram) map[string]map[string]string` | (currentState, event) → nextState (`[*]` 초기 전이 제외) |
| `renderStateFile` | `(id, transMap) string` | Go source 조립 (transitions var + CanTransition + NextState) |
| `renderTransitionEntries` | `(transMap) string` | map literal 라인 |
| `renderCanTransitionFile` / `renderNextStateFile` | `(id, transMap) string` | 함수 본문 |
| `writeStateFile` | `(dir, id, source) error` | `os.WriteFile` |

## 산출물

```
arts/backend/internal/statemachine/
├── workflow.go           ← StateDiagram.ID="Workflow"
├── reservation.go        ← StateDiagram.ID="Reservation"
└── ...
```

다이어그램 1 개 = 파일 1 개. 모든 파일 `package statemachine`. 함수/변수명에 ID 접두사로 충돌 방지.

## 네이밍 규약

| stateDiagram ID | 파일명 | Transitions | CanTransition | NextState |
|---|---|---|---|---|
| `Workflow` | `workflow.go` | `WorkflowTransitions` | `WorkflowCanTransition` | `WorkflowNextState` |
| `Reservation` | `reservation.go` | `ReservationTransitions` | `ReservationCanTransition` | `ReservationNextState` |

파일명: ID → snake_case. 함수/변수명: ID + PascalCase suffix.

## 핸들러 사용 예

SSaC `@state Workflow {status: wf.Status} "ActivateWorkflow" "Cannot activate" 409` →
```go
if !statemachine.WorkflowCanTransition(wf.Status, "ActivateWorkflow") {
    return api.ActivateWorkflow409JSONResponse{Error: strPtr("Cannot activate")}, nil
}
```
import: `"<module>/internal/statemachine"`.

## 정책

- `[*] --> draft` (초기 전이) 는 맵에 포함하지 않음. DDL `DEFAULT 'draft'` 가 보장 (XDM-28).
- self-transition (`active --> active: ExecuteWorkflow`) 은 정상 등록 — `CanTransition` true 반환, 상태 값 불변 (no-op put 또는 put 생략).
- 외부 import 없음 (표준 라이브러리도 불필요).
