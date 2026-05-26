//ff:type feature=gen-ir type=model
//ff:what DeleteOp -- @delete 시퀀스의 IR 표현 (DB DELETE, 결과 없음)

package ir

// DeleteOp represents a @delete sequence: a database delete that does not bind
// a result variable (void return).
type DeleteOp struct {
	// Model is the sqlc table/model name.
	Model string

	// Method is the sqlc query method name (e.g. "Delete").
	Method string

	// Args are the delete arguments.
	Args []FieldArg
}
