//ff:func feature=statemachine type=util control=sequence topic=state-transition
//ff:what WorkflowNextState — 현재 상태 + 이벤트 → 다음 상태 반환 (불가 시 "")
//ff:checked llm=yongol-gen hash=80ca259a
package statemachine

// WorkflowNextState returns the target state after a valid transition.
// Returns empty string when the transition is not allowed.
func WorkflowNextState(currentState, event string) string {
	events, ok := WorkflowTransitions[currentState]
	if !ok {
		return ""
	}
	return events[event]
}
