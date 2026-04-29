//ff:func feature=gen-gogin type=util control=sequence
//ff:what jsonb — JSONB/JSON 컬럼의 map[string]any 매핑 (3 nullability)

package types

// jsonbBinding returns the binding for a JSONB / JSON column.
//
//   - NOT NULL              → map[string]any
//   - NULLABLE              → *map[string]any
//   - NOT NULL DEFAULT '{}' → caller wraps DefaultLiteral as []byte("'{}'")
//
// sqlc pgx/v5 emits []byte for json / jsonb without an override. yongol
// wraps that to map[string]any for convenience and to match OpenAPI
// `type: object` semantics. The convert site unmarshals via json.Unmarshal
// (handled by the caller because the error must propagate as a 500).
func jsonbBinding(notNull bool, defaultLiteral string) GoTypeBinding {
	apiField := "map[string]interface{}"
	if !notNull {
		apiField = "*map[string]interface{}"
	}
	return GoTypeBinding{
		SqlcGoType:     "[]byte",
		ApiField:       apiField,
		Imports:        []string{"encoding/json"},
		ConvertExpr:    "{row}.{field}",
		InsertExpr:     "{var}",
		ResponseExpr:   "{var}.{field}",
		DefaultLiteral: defaultLiteral,
		Kind:           KindJSONB,
		Supported:      true,
	}
}
