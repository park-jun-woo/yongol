//ff:func feature=gen-gogin type=util control=sequence
//ff:what pgtypeInet — INET/CIDR 컬럼의 pgtype.Inet 매핑 (NeedsOverride=true)

package types

// pgtypeInet returns the binding for an INET / CIDR column. sqlc pgx/v5
// requires an explicit override for `inet` and `cidr` because the
// standard library does not provide a 1-1 type and pgx exposes the value
// via netip / netaddr-style structs.
//
// The api surface uses string (CIDR notation, e.g. "10.0.0.1/24") which
// matches PostgreSQL's textual output and OpenAPI's `format: cidr` /
// `format: ipv4` conventions.
func pgtypeInet(notNull bool, defaultLiteral string) GoTypeBinding {
	apiField := "string"
	if !notNull {
		apiField = "*string"
	}
	return GoTypeBinding{
		SqlcGoType:    "pgtype.Inet",
		NeedsOverride: true,
		ApiField:      apiField,
		Imports: []string{
			"github.com/jackc/pgx/v5/pgtype",
		},
		ConvertExpr:    "pgInetToString({row}.{field})",
		InsertExpr:     "{var}",
		ResponseExpr:   "pgInetToString({var}.{field})",
		DefaultLiteral: defaultLiteral,
		Kind:           KindPgtype,
		Supported:      true,
	}
}
