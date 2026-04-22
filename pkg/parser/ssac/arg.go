//ff:type feature=ssac-parse type=model
//ff:what Arg — type representing a single function call argument
package ssac

// Arg represents a single function call argument.
type Arg struct {
	Source   string // "request", a variable name, or "" (literal)
	Field    string // "CourseID", "ID", etc.
	Literal  string // raw literal text (quotes stripped for quoted strings)
	IsQuoted bool   // whether Literal was a "..." quoted string — used for type inference
}
