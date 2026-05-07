//ff:func feature=gen-gogin type=util control=sequence
//ff:what nativeFloat — REAL/FLOAT/FLOAT4/FLOAT8 매핑 (NOT NULL → float64, NULL → pgtype.Float8)

package types

// nativeFloat returns the binding for a float-family column.
// sqlc pgx/v5 emits pgtype.Float8 for nullable double precision and
// pgtype.Float4 for nullable real/float4 — pgtypex bridge.
func nativeFloat(notNull bool, defaultLiteral string) GoTypeBinding {
	if notNull {
		return GoTypeBinding{
			SqlcGoType:     "float64",
			ApiField:       "float64",
			ConvertExpr:    "{row}.{field}",
			InsertExpr:     "{var}",
			ResponseExpr:   "{var}.{field}",
			DefaultLiteral: defaultLiteral,
			Kind:           KindNative,
			Supported:      true,
		}
	}
	return pgtypeFloat8(defaultLiteral)
}
