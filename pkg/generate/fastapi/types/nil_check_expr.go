//ff:func feature=gen-fastapi type=util control=sequence
//ff:what nilCheckExpr — NotNull 여부에 따라 Python None 검사 표현식 결정

package types

// nilCheckExpr returns the nil-check predicate template. NOT NULL columns
// return "" (caller uses language-native zero comparison). NULLABLE columns
// return "{var} is None".
func nilCheckExpr(notNull bool) string {
	if notNull {
		return ""
	}
	return "{var} is None"
}
