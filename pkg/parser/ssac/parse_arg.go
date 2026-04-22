//ff:func feature=ssac-parse type=util control=sequence
//ff:what parseArg — parses a single argument and returns an Arg
package ssac

import "strings"

// parseArg parses a single argument.
func parseArg(s string) Arg {
	s = strings.TrimSpace(s)
	// "literal" — mark quoted strings with IsQuoted=true for int/string disambiguation.
	if strings.HasPrefix(s, `"`) && strings.HasSuffix(s, `"`) {
		return Arg{Literal: s[1 : len(s)-1], IsQuoted: true}
	}
	// numeric, boolean, nil literal — check before the dot test (so 3.14 is not parsed as source.Field)
	if IsLiteral(s) {
		return Arg{Literal: s}
	}
	// source.Field
	dotIdx := strings.IndexByte(s, '.')
	if dotIdx > 0 {
		return Arg{Source: s[:dotIdx], Field: s[dotIdx+1:]}
	}
	// bare variable (shouldn't happen in valid syntax, but handle gracefully)
	return Arg{Source: s}
}
