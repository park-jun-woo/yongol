//ff:func feature=validate type=util control=sequence topic=ssac-structural
//ff:what lcFirstEval — PascalCase → camelCase for @eval Model lookups

package ssac

// lcFirstEval lowercases the first byte of an ASCII Method name so that an
// @eval Model written in PascalCase (e.g. "IsZeroBalance") matches FuncSpec
// names which are camelCase (`isZeroBalance`). Non-ASCII inputs pass through.
func lcFirstEval(s string) string {
	if s == "" {
		return s
	}
	c := s[0]
	if c >= 'A' && c <= 'Z' {
		return string(c+32) + s[1:]
	}
	return s
}
