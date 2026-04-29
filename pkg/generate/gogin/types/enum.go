//ff:func feature=gen-gogin type=util control=sequence
//ff:what enumBinding — VARCHAR(N) + CHECK IN (...) 컬럼의 string + apiCast 매핑

package types

// enumBinding returns the binding for a CHECK IN (...) enum column. The
// underlying PG storage is VARCHAR(N) so sqlc emits string / *string;
// yongol layers an apiCast on top so the api struct surface uses the
// generated named type (e.g. api.OrderStatus). The cast itself is
// applied by the emit site because it depends on the api type name,
// which is per-column and not derivable from the binding alone.
//
// Caller signals enum-ness by passing checkEnum != nil from the parsed
// Column. The slice contents are not consumed here — only its non-empty
// presence matters; the values surface through validate to the api enum.
func enumBinding(notNull bool, defaultLiteral string) GoTypeBinding {
	apiField := "string"
	if !notNull {
		apiField = "*string"
	}
	return GoTypeBinding{
		SqlcGoType:     apiField, // string or *string — same as native
		ApiField:       apiField,
		ConvertExpr:    "{row}.{field}",
		InsertExpr:     "{var}",
		ResponseExpr:   "{var}.{field}",
		DefaultLiteral: defaultLiteral,
		Kind:           KindEnum,
		Supported:      true,
	}
}
