//ff:type feature=gen-ir type=model
//ff:what EmptyOp -- @empty 시퀀스의 IR 표현 (nil 가드 → 에러 분기)

package ir

// EmptyOp represents an @empty sequence: a nil guard that returns an error
// response when the target variable is nil (e.g. 404 Not Found).
type EmptyOp struct {
	// VarName is the variable to nil-check (e.g. "course").
	VarName string

	// Message is the error message returned when the guard triggers.
	Message string

	// StatusCode is the HTTP status code (e.g. 404).
	StatusCode int
}
