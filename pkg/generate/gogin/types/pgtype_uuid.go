//ff:func feature=gen-gogin type=util control=sequence
//ff:what pgtypeUUID — UUID 컬럼의 pgtype.UUID 매핑 (NeedsOverride=true, pgtypex bridge)

package types

// pgtypeUUID returns the binding for a UUID column. sqlc pgx/v5 has no
// default mapping for PG `uuid`; without an explicit override it would
// emit `interface{}` or refuse to compile. Q-12 enforces the matching
// sqlc.yaml override on the user side.
//
// The convert/insert sites route through ssac/pkg/pgtypex which
// centralises the pgtype.UUID ↔ openapi_types.UUID bridging.
func pgtypeUUID(notNull bool, defaultLiteral string) GoTypeBinding {
	apiField := "openapi_types.UUID"
	if !notNull {
		apiField = "*openapi_types.UUID"
	}
	toFunc, fromFunc := "pgtypex.ToPgUUID", "pgtypex.FromPgUUID"
	if !notNull {
		toFunc, fromFunc = "pgtypex.ToPgUUIDPtr", "pgtypex.FromPgUUIDPtr"
	}
	return GoTypeBinding{
		SqlcGoType:    "pgtype.UUID",
		NeedsOverride: true,
		ApiField:      apiField,
		Imports: []string{
			"github.com/jackc/pgx/v5/pgtype",
			"github.com/oapi-codegen/runtime/types",
			"github.com/park-jun-woo/ssac/pkg/pgtypex",
		},
		ConvertExpr:    fromFunc + "({row}.{field})",
		InsertExpr:     toFunc + "({var})",
		ResponseExpr:   fromFunc + "({var}.{field})",
		NilCheckExpr:   "pgtypex.IsNilPgUUID({var})",
		DefaultLiteral: defaultLiteral,
		Kind:           KindPgtype,
		Supported:      true,
	}
}
