//ff:type feature=gen-ir type=model
//ff:what FieldArg -- SSaC 인자의 통합 IR 표현 (positional + map key 겸용)

package ir

// ParamLocation classifies where a request field originates in the HTTP
// request. Mirrors the classification that gogin/ssac/methodGen performs
// via addParam (path/query) and mapRequestValue (body fallback).
type ParamLocation string

const (
	LocPath    ParamLocation = "path"    // OpenAPI path parameter
	LocQuery   ParamLocation = "query"   // OpenAPI query parameter
	LocBody    ParamLocation = "body"    // OpenAPI request body property
	LocVar     ParamLocation = "var"     // previous sequence result variable
	LocLiteral ParamLocation = "literal" // inline literal value
	LocUser    ParamLocation = "user"    // currentUser field
)

// FieldArg is the unified IR representation for SSaC arguments. It merges the
// two parser-level representations (ssac.Arg for positional args and
// map[string]string for state/auth/publish inputs) into a single struct.
type FieldArg struct {
	// Key is the map key for state/auth/publish inputs.
	// Empty string for positional arguments (get/post/put/delete/call).
	Key string

	// Source is "request", a variable name, or empty (for literals).
	Source string

	// Field is the field accessor (e.g. ".ID", ".Status").
	Field string

	// Literal holds a raw literal value when the argument is not a reference.
	Literal string

	// IsQuoted is true when the literal was a "..." quoted string.
	IsQuoted bool

	// Location classifies where a request field originates (path, query,
	// body, var, literal, user). Populated by enrichFieldArgLocations from
	// OpenAPI parameter metadata. Empty string for non-request sources
	// when OpenAPI doc is absent.
	Location ParamLocation

	// ColumnName is the DDL snake_case column name corresponding to this
	// argument's Key, resolved via PascalToSnake matching against DDL
	// table columns. Empty when no DDL match is found.
	ColumnName string

	// IsPK is true when ColumnName matches a DDL PrimaryKey column.
	IsPK bool
}
