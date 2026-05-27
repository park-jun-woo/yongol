//ff:type feature=gen-ir type=model
//ff:what GetOp -- @get 시퀀스의 IR 표현 (DB 조회 + 결과 바인딩)

package ir

// GetOp represents an @get sequence: a database read that binds a result
// variable.
type GetOp struct {
	// VarName is the result variable name (e.g. "course").
	VarName string

	// VarType is the model type name (e.g. "Course").
	VarType string

	// Model is the sqlc table/model name.
	Model string

	// Method is the sqlc query method name (e.g. "FindByID").
	Method string

	// Args are the query arguments.
	Args []FieldArg

	// IsList is true when the query returns a slice.
	IsList bool

	// FollowedByGuard is the OpKind of the immediately following @empty or
	// @exists guard that targets this GetOp's VarName. When set (OpEmpty or
	// OpExists), the Go renderer emits errors.Is(err, pgx.ErrNoRows) instead
	// of a plain error propagation so the guard logic can check the zero
	// value. Zero value (OpGet) means no guard follows.
	FollowedByGuard OpKind

	// IsCount is true when the result type is a scalar integer (int64, int32,
	// int), indicating a COUNT query rather than a row query.
	IsCount bool

	// PaginationArgs holds pagination-specific arguments (cursor, per_page,
	// page_offset, page, limit, offset) separated from the where-clause Args.
	// Renderers targeting ORMs (Prisma, SQLAlchemy) use this to emit
	// framework-native pagination calls.
	PaginationArgs []FieldArg
}
