//ff:func feature=gen-gogin type=util control=sequence
//ff:what pointerInteger — NULLABLE 정수 컬럼의 *int64 매핑

package types

// pointerInteger produces the GoTypeBinding for a nullable integer
// column. sqlc pgx/v5 emits *int64 for nullable BIGINT/INT, so SqlcGoType
// also carries the *T form and NeedsOverride stays false.
func pointerInteger(defaultLiteral string) GoTypeBinding {
	return GoTypeBinding{
		SqlcGoType:     "*int64",
		ApiField:       "*int64",
		ConvertExpr:    "{row}.{field}",
		InsertExpr:     "{var}",
		ResponseExpr:   "{var}.{field}",
		DefaultLiteral: defaultLiteral,
		Kind:           KindPointer,
		Supported:      true,
	}
}
