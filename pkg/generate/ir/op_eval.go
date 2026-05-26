//ff:type feature=gen-ir type=model
//ff:what EvalOp -- @eval 시퀀스의 IR 표현 (bool 함수 평가 → 에러 분기)

package ir

// EvalOp represents an @eval sequence: a boolean function evaluation that
// returns an error response when the result is true.
type EvalOp struct {
	// Package is the package prefix. Empty for same-package.
	Package string

	// Function is the function name (e.g. "IsExpired").
	Function string

	// Args are the function arguments.
	Args []FieldArg

	// Message is the error message returned when the evaluation is true.
	Message string

	// StatusCode is the HTTP status code on failure (default 400).
	StatusCode int
}
