//ff:func feature=migration type=parser control=selection
//ff:what dispatchColumnAttr — rest[i] 키워드에 따라 컬럼 속성 핸들러 호출 + 소비한 토큰 수 반환
package migration

import "strings"

// dispatchColumnAttr inspects rest[i] and delegates to the appropriate
// sub-handler. Returns the number of tokens consumed (1 when the token
// was skipped unchanged).
func dispatchColumnAttr(t *Table, col *Column, rest []string, i int) int {
	tok := strings.ToUpper(rest[i])
	switch {
	case tok == "NOT" && i+1 < len(rest) && strings.ToUpper(rest[i+1]) == "NULL":
		col.Nullable = false
		return 2
	case tok == "NULL":
		col.Nullable = true
		return 1
	case tok == "DEFAULT":
		def, consumed := collectDefaultExpr(rest[i+1:])
		col.Default = NormalizeDefault(def)
		return 1 + consumed
	case tok == "PRIMARY" && i+1 < len(rest) && strings.ToUpper(rest[i+1]) == "KEY":
		t.PrimaryKey = []string{col.Name}
		col.Nullable = false
		return 2
	case tok == "UNIQUE":
		t.Indexes = append(t.Indexes, &Index{
			Name:    UniqueName(t.Name, []string{col.Name}),
			Columns: []string{col.Name},
			Unique:  true,
		})
		return 1
	case tok == "REFERENCES":
		fk, consumed := parseInlineRef(t, col.Name, rest[i+1:])
		if fk != nil {
			t.ForeignKeys = append(t.ForeignKeys, fk)
		}
		return 1 + consumed
	case tok == "CHECK":
		return applyInlineCheck(t, col, rest, i)
	}
	return 1
}
