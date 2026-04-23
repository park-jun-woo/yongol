//ff:func feature=migration type=parser control=iteration dimension=1
//ff:what parseColumnList — "(a, b, c)" 또는 "a, b, c" 를 컬럼명 슬라이스로
package migration

import "strings"

// parseColumnList parses a comma-separated list of column identifiers,
// optionally wrapped in parentheses, and canonicalises each ident.
func parseColumnList(s string) []string {
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "(") && strings.HasSuffix(s, ")") {
		s = s[1 : len(s)-1]
	}
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		out = append(out, canonIdent(p))
	}
	return out
}
