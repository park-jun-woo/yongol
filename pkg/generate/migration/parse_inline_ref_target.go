//ff:func feature=migration type=parser control=sequence
//ff:what parseInlineRefTarget — REFERENCES 뒤 첫 토큰(들)에서 refTable/refCol 추출
package migration

import "strings"

// parseInlineRefTarget reads the target table and optional (column) of a
// `REFERENCES ...` clause. Returns (refTable, refCol, consumed).
func parseInlineRefTarget(toks []string) (string, string, int) {
	target := toks[0]
	if p := strings.Index(target, "("); p >= 0 {
		refTable := canonIdent(target[:p])
		refCol := ""
		if end := strings.LastIndex(target, ")"); end > p {
			refCol = canonIdent(strings.TrimSpace(target[p+1 : end]))
		}
		return refTable, refCol, 1
	}
	if len(toks) > 1 && strings.HasPrefix(toks[1], "(") {
		return canonIdent(target), canonIdent(innerParens(toks[1])), 2
	}
	return canonIdent(target), "", 1
}
