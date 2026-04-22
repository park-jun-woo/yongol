//ff:func feature=ssac-parse type=util control=sequence
//ff:what 단일 인자를 파싱하여 Arg 반환
package ssac

import "strings"

// parseArg는 단일 인자를 파싱한다.
func parseArg(s string) Arg {
	s = strings.TrimSpace(s)
	// "literal" — quoted string 은 IsQuoted=true 로 표시하여 int/string 구분 가능하게.
	if strings.HasPrefix(s, `"`) && strings.HasSuffix(s, `"`) {
		return Arg{Literal: s[1 : len(s)-1], IsQuoted: true}
	}
	// numeric, boolean, nil literal — dot 검사보다 먼저 (3.14가 source.Field로 파싱되지 않도록)
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
