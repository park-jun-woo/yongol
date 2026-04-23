//ff:func feature=migration type=parser control=sequence
//ff:what parseTableFKRef — 테이블 FK 의 REFERENCES 뒤 target 파싱 → (refTable, refCols, consumed)
package migration

import "strings"

// parseTableFKRef resolves the target table / columns of a table-scoped
// FOREIGN KEY clause. `toks` is the full tokenized tail starting with
// "(localCols) REFERENCES target ...".
func parseTableFKRef(toks []string) (string, []string, int) {
	target := toks[2]
	consumed := 3
	if p := strings.Index(target, "("); p >= 0 {
		refTable := canonIdent(target[:p])
		var refCols []string
		if end := strings.LastIndex(target, ")"); end > p {
			refCols = parseColumnList(target[p+1 : end])
		}
		return refTable, refCols, consumed
	}
	if consumed < len(toks) && strings.HasPrefix(toks[consumed], "(") {
		return canonIdent(target), parseColumnList(innerParensFull(toks[consumed])), consumed + 1
	}
	return canonIdent(target), nil, consumed
}
