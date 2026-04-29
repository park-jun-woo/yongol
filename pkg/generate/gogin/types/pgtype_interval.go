//ff:func feature=gen-gogin type=util control=sequence
//ff:what pgtypeInterval — INTERVAL 컬럼의 pgtype.Interval 매핑 (NeedsOverride=true)

package types

// pgtypeInterval returns the binding for an INTERVAL column. sqlc pgx/v5
// has no default mapping; pgtype.Interval exposes Microseconds + Days +
// Months. The api surface uses string (ISO 8601 duration "PT1H30M"
// convention) — the convert helper emits that form.
func pgtypeInterval(notNull bool, defaultLiteral string) GoTypeBinding {
	apiField := "string"
	if !notNull {
		apiField = "*string"
	}
	return GoTypeBinding{
		SqlcGoType:    "pgtype.Interval",
		NeedsOverride: true,
		ApiField:      apiField,
		Imports: []string{
			"github.com/jackc/pgx/v5/pgtype",
		},
		ConvertExpr:    "pgIntervalToString({row}.{field})",
		InsertExpr:     "{var}",
		ResponseExpr:   "pgIntervalToString({var}.{field})",
		DefaultLiteral: defaultLiteral,
		Kind:           KindPgtype,
		Supported:      true,
	}
}
