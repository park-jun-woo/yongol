//ff:func feature=gen-nestjs type=util control=sequence
//ff:what nullableAPIType — NotNull 여부에 따라 TypeScript 타입에 " | null" 접미사 추가

package types

// nullableAPIType appends " | null" to the base TypeScript type when the
// column is nullable (notNull=false). For NOT NULL columns the base type
// is returned as-is.
func nullableAPIType(base string, notNull bool) string {
	if notNull {
		return base
	}
	return base + " | null"
}
