//ff:type feature=gen-ir type=model
//ff:what CallOp -- @call 시퀀스의 IR 표현 (외부 함수 호출 + 선택적 결과 바인딩)

package ir

// CallOp represents a @call sequence: an external function invocation with
// optional result binding and error handling.
type CallOp struct {
	// Package is the package prefix (e.g. "billing"). Empty for same-package.
	Package string

	// TargetFeature is the feature folder name of the call target (e.g.
	// "webhookdelivery"). Derived from Package via lowercase convention.
	// Empty for same-package calls. Used by module renderers to add
	// cross-feature imports.
	TargetFeature string

	// Function is the function name (e.g. "HoldEscrow").
	Function string

	// Args are the function arguments.
	Args []FieldArg

	// ResultVar is the result variable name. Empty when the call has no result
	// binding.
	ResultVar string

	// ResultType is the result type name. Empty when no result binding.
	ResultType string

	// ErrStatus is the HTTP status code for call errors (default 500).
	ErrStatus int

	// Message is the error message returned on call failure.
	Message string
}
