//ff:func feature=gen-gogin type=util control=selection
//ff:what nativeBoolean — BOOLEAN/BOOL 매핑 (NOT NULL → bool, NULL → pgtype.Bool)

package types

// nativeBoolean returns the binding for a boolean column.
// sqlc pgx/v5 emits pgtype.Bool for nullable boolean columns — pgtypex bridge.
func nativeBoolean(notNull bool, defaultLiteral string) GoTypeBinding {
	if notNull {
		return GoTypeBinding{
			SqlcGoType:     "bool",
			ApiField:       "bool",
			ConvertExpr:    "{row}.{field}",
			InsertExpr:     "{var}",
			ResponseExpr:   "{var}.{field}",
			DefaultLiteral: defaultLiteral,
			Kind:           KindNative,
			Supported:      true,
		}
	}
	return pgtypeBool(defaultLiteral)
}
