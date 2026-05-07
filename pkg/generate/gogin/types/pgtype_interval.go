//ff:func feature=gen-gogin type=util control=sequence
//ff:what pgtypeInterval — INTERVAL 컬럼의 pgtype.Interval 매핑 (NeedsOverride=true, pgtypex bridge)

package types

// pgtypeInterval returns the binding for an INTERVAL column. sqlc pgx/v5
// has no default mapping; pgtype.Interval exposes Microseconds + Days +
// Months. The api surface uses string (ISO 8601 duration "PT1H30M"
// convention).
func pgtypeInterval(notNull bool, defaultLiteral string) GoTypeBinding {
	apiField := "string"
	if !notNull {
		apiField = "*string"
	}
	toFunc, fromFunc := "pgtypex.ToPgInterval", "pgtypex.FromPgInterval"
	if !notNull {
		toFunc, fromFunc = "pgtypex.ToPgIntervalPtr", "pgtypex.FromPgIntervalPtr"
	}
	return GoTypeBinding{
		SqlcGoType:    "pgtype.Interval",
		NeedsOverride: true,
		ApiField:      apiField,
		Imports: []string{
			"github.com/jackc/pgx/v5/pgtype",
			"github.com/park-jun-woo/ssac/pkg/pgtypex",
		},
		ConvertExpr:    fromFunc + "({row}.{field})",
		InsertExpr:     toFunc + "({var})",
		ResponseExpr:   fromFunc + "({var}.{field})",
		NilCheckExpr:   "pgtypex.IsNilPgInterval({var})",
		DefaultLiteral: defaultLiteral,
		Kind:           KindPgtype,
		Supported:      true,
	}
}
