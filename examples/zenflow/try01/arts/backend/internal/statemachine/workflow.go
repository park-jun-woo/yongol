//ff:type feature=statemachine type=model
//ff:what WorkflowTransitions — 상태 전이 테이블 (currentState, event → nextState)
package statemachine

// WorkflowTransitions maps (currentState, event) → nextState.
// Generated from states/workflow.md — do not edit.
var WorkflowTransitions = map[string]map[string]string{
	"active": {"ExecuteWorkflow": "active"},
	"draft": {"ActivateWorkflow": "active"},
	"paused": {"ActivateWorkflow": "active"},
}
