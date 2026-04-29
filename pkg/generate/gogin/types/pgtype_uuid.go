//ff:func feature=gen-gogin type=util control=sequence
//ff:what pgtypeUUID — UUID 컬럼의 pgtype.UUID 매핑 (NeedsOverride=true)

package types

// pgtypeUUID returns the binding for a UUID column. sqlc pgx/v5 has no
// default mapping for PG `uuid`; without an explicit override it would
// emit `interface{}` or refuse to compile. Q-12 enforces the matching
// sqlc.yaml override on the user side.
//
// The convert site routes through pgUUIDToString (emitted into
// internal/service/pg_uuid_to_string.go) which centralises the
// `Valid + [16]byte → canonical UUID string` extraction.
func pgtypeUUID(notNull bool, defaultLiteral string) GoTypeBinding {
	apiField := "openapi_types.UUID"
	if !notNull {
		apiField = "*openapi_types.UUID"
	}
	return GoTypeBinding{
		SqlcGoType:     "pgtype.UUID",
		NeedsOverride:  true,
		ApiField:       apiField,
		Imports: []string{
			"github.com/jackc/pgx/v5/pgtype",
			"github.com/oapi-codegen/runtime/types",
		},
		ConvertExpr:    "pgUUIDToString({row}.{field})",
		InsertExpr:     "{var}",
		ResponseExpr:   "pgUUIDToString({var}.{field})",
		DefaultLiteral: defaultLiteral,
		Kind:           KindPgtype,
		Supported:      true,
	}
}
