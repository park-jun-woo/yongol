//ff:func feature=gen-gogin type=util control=sequence
//ff:what nativeString — VARCHAR/TEXT/CHAR/BPCHAR 매핑 (NOT NULL → string, NULL → pgtype.Text)

package types

// nativeString returns the binding for a string-family column. The
// parameter list (e.g. "255" for VARCHAR(255)) is not encoded into the
// binding itself; length validation lives in the OpenAPI / DDL layer.
// sqlc pgx/v5 emits pgtype.Text for nullable text columns — pgtypex bridge.
func nativeString(notNull bool, defaultLiteral string) GoTypeBinding {
	if notNull {
		return GoTypeBinding{
			SqlcGoType:     "string",
			ApiField:       "string",
			ConvertExpr:    "{row}.{field}",
			InsertExpr:     "{var}",
			ResponseExpr:   "{var}.{field}",
			DefaultLiteral: defaultLiteral,
			Kind:           KindNative,
			Supported:      true,
		}
	}
	return pgtypeText(defaultLiteral)
}
