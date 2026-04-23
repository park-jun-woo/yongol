//ff:func feature=migration type=util control=sequence
//ff:what splitTypeParams — "VARCHAR(255)" / "NUMERIC(10,2)" 에서 base 와 파라미터 분리
package migration

import "strings"

// splitTypeParams splits a type string into (base, params) where params
// is the comma-separated list inside parens ("" if no params).
func splitTypeParams(s string) (string, string) {
	i := strings.Index(s, "(")
	if i < 0 || !strings.HasSuffix(s, ")") {
		return s, ""
	}
	return strings.TrimSpace(s[:i]), s[i+1 : len(s)-1]
}
