//ff:type feature=gen-ir type=model
//ff:what PutOp -- @put 시퀀스의 IR 표현 (DB UPDATE, 결과 없음)

package ir

// PutOp represents an @put sequence: a database update that does not bind a
// result variable (void return).
type PutOp struct {
	// Model is the sqlc table/model name.
	Model string

	// Method is the sqlc query method name (e.g. "UpdateStatus").
	Method string

	// Args are the update arguments.
	Args []FieldArg
}
