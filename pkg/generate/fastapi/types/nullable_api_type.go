//ff:func feature=gen-fastapi type=util control=sequence
//ff:what nullableAPIType — NotNull 여부에 따라 Python 타입에 " | None" 접미사 추가

package types

// nullableAPIType appends " | None" to the base Python type when the
// column is nullable (notNull=false). For NOT NULL columns the base type
// is returned as-is. Uses Python 3.10+ union syntax.
func nullableAPIType(base string, notNull bool) string {
	if notNull {
		return base
	}
	return base + " | None"
}
