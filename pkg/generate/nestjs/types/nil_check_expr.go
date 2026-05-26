//ff:func feature=gen-nestjs type=util control=sequence
//ff:what nilCheckExpr — NotNull 여부에 따라 TypeScript null 검사 표현식 결정

package types

// nilCheckExpr returns the nil-check predicate template. NOT NULL columns
// return "" (caller uses language-native zero comparison). NULLABLE columns
// return "{var} === null".
func nilCheckExpr(notNull bool) string {
	if notNull {
		return ""
	}
	return "{var} === null"
}
