//ff:type feature=gen-ir type=model
//ff:what ExistsOp -- @exists 시퀀스의 IR 표현 (non-nil 가드 → 에러 분기)

package ir

// ExistsOp represents an @exists sequence: a non-nil guard that returns an
// error response when the target variable is NOT nil (e.g. 409 Conflict).
type ExistsOp struct {
	// VarName is the variable to check (e.g. "existing").
	VarName string

	// Message is the error message returned when the guard triggers.
	Message string

	// StatusCode is the HTTP status code (e.g. 409).
	StatusCode int
}
