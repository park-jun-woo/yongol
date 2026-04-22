//ff:type feature=statemachine type=model topic=states
//ff:what 상태 전이를 나타내는 구조체
package statemachine

// Transition represents a single state transition.
type Transition struct {
	From  string // source state
	To    string // target state
	Event string // operationId / SSaC function name
	Line  int    // 1-based line number of the transition in the source file (0 = unknown)
}
