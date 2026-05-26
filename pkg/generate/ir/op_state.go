//ff:type feature=gen-ir type=model
//ff:what StateOp -- @state 시퀀스의 IR 표현 (Mermaid stateDiagram 전이 검증)

package ir

// StateOp represents a @state sequence: a state machine transition validation
// against a Mermaid stateDiagram.
type StateOp struct {
	// Diagram is the state diagram identifier (e.g. "reservation").
	Diagram string

	// Inputs are the key-value pairs providing state context.
	Inputs []FieldArg

	// Transition is the transition name to validate (e.g. "cancel").
	Transition string

	// Message is the error message returned on invalid transition.
	Message string

	// StatusCode is the HTTP status code on failure (default 409).
	StatusCode int
}
