//ff:func feature=gen-gogin type=util control=sequence
//ff:what byteaBinding — BYTEA 컬럼의 []byte 매핑 (slice 자체가 nullable)

package types

// byteaBinding returns the binding for a BYTEA column. sqlc pgx/v5 emits
// []byte for both nullable and NOT NULL columns; the slice itself is
// nullable so no pointer wrapping is required. JSON encoding produces
// base64 by default via the standard library — no custom marshalling
// needed.
func byteaBinding(notNull bool, defaultLiteral string) GoTypeBinding {
	_ = notNull // kept for symmetry with the other constructors
	return GoTypeBinding{
		SqlcGoType:     "[]byte",
		ApiField:       "[]byte",
		ConvertExpr:    "{row}.{field}",
		InsertExpr:     "{var}",
		ResponseExpr:   "{var}.{field}",
		DefaultLiteral: defaultLiteral,
		Kind:           KindBytea,
		Supported:      true,
	}
}
