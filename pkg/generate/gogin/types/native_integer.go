//ff:func feature=gen-gogin type=util control=sequence
//ff:what nativeInteger — BIGINT/INT/SMALLINT 계열 매핑 (NOT NULL → int64, NULL → pgtype.Int8)

package types

// nativeInteger returns the binding for an integer-family column.
// notNull selects the NOT NULL (Native) vs NULLABLE (pgtype) form.
// sqlc pgx/v5 emits pgtype.Int8 for nullable BIGINT — pgtypex bridge.
func nativeInteger(notNull bool, defaultLiteral string) GoTypeBinding {
	if notNull {
		return GoTypeBinding{
			SqlcGoType:     "int64",
			ApiField:       "int64",
			ConvertExpr:    "{row}.{field}",
			InsertExpr:     "{var}",
			ResponseExpr:   "{var}.{field}",
			DefaultLiteral: defaultLiteral,
			Kind:           KindNative,
			Supported:      true,
		}
	}
	return pgtypeInt8(defaultLiteral)
}
