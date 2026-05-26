//ff:type feature=gen-ir type=model
//ff:what FieldArg -- SSaC 인자의 통합 IR 표현 (positional + map key 겸용)

package ir

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
}
