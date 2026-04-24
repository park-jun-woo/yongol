//ff:func feature=statemachine type=util control=sequence topic=state-transition
//ff:what WorkflowCanTransition — 현재 상태에서 이벤트 전이 가능 여부 반환
//ff:checked llm=yongol-gen hash=0fffc0ba
package statemachine

// WorkflowCanTransition returns true when event is a valid transition
// from currentState.
func WorkflowCanTransition(currentState, event string) bool {
	events, ok := WorkflowTransitions[currentState]
	if !ok {
		return false
	}
	_, ok = events[event]
	return ok
}
