//ff:type feature=gen-ir type=model
//ff:what ResponseOp -- @response 시퀀스의 IR 표현 (JSON 응답 조립)

package ir

// ResponseOp represents a @response sequence: assembling the final JSON
// response body from named fields or a single variable.
type ResponseOp struct {
	// Fields maps JSON field names to their source variables/expressions.
	// Non-nil when the response is a multi-field object.
	Fields []ResponseField

	// SingleVar is the variable name when the response is a direct single-
	// variable return (e.g. "course"). Empty when Fields is used.
	SingleVar string
}
