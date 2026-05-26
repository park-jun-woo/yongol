//ff:type feature=gen-ir type=model
//ff:what PostOp -- @post 시퀀스의 IR 표현 (DB INSERT + 결과 바인딩)

package ir

// PostOp represents an @post sequence: a database insert that binds a result
// variable.
type PostOp struct {
	// VarName is the result variable name (e.g. "course").
	VarName string

	// VarType is the model type name (e.g. "Course").
	VarType string

	// Model is the sqlc table/model name.
	Model string

	// Method is the sqlc query method name (e.g. "Create").
	Method string

	// Args are the insert arguments.
	Args []FieldArg

	// IsList is true when the query returns a slice.
	IsList bool
}
