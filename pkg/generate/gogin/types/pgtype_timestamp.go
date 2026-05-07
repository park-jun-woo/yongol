//ff:func feature=gen-gogin type=util control=selection
//ff:what pgtypeTimestamp — TIMESTAMPTZ/TIMESTAMP/DATE 컬럼의 pgtype 매핑 (pgtypex bridge)

package types

// pgtypeTimestamp returns the binding for a temporal column. sqlc pgx/v5
// emits pgtype.Timestamptz / pgtype.Timestamp / pgtype.Date depending on
// the source PG type.
//
// NeedsOverride is true because the default sqlc map for `timestamp /
// timestamptz / date` predates pgx/v5 in some configurations and the
// project standardises on pgtype.Timestamptz to round-trip TZ correctly.
func pgtypeTimestamp(head string, notNull bool, defaultLiteral string) GoTypeBinding {
	sqlcType := "pgtype.Timestamptz"
	suffix := "Timestamptz"
	switch head {
	case "TIMESTAMP":
		sqlcType = "pgtype.Timestamp"
		suffix = "Timestamp"
	case "DATE":
		sqlcType = "pgtype.Date"
		suffix = "Date"
	}
	apiField := "time.Time"
	if !notNull {
		apiField = "*time.Time"
	}
	toFunc := "pgtypex.ToPg" + suffix
	fromFunc := "pgtypex.FromPg" + suffix
	if !notNull {
		toFunc += "Ptr"
		fromFunc += "Ptr"
	}
	return GoTypeBinding{
		SqlcGoType:    sqlcType,
		NeedsOverride: true,
		ApiField:      apiField,
		Imports: []string{
			"github.com/jackc/pgx/v5/pgtype",
			"time",
			"github.com/park-jun-woo/ssac/pkg/pgtypex",
		},
		ConvertExpr:    fromFunc + "({row}.{field})",
		InsertExpr:     toFunc + "({var})",
		ResponseExpr:   fromFunc + "({var}.{field})",
		NilCheckExpr:   "pgtypex.IsNilPg" + suffix + "({var})",
		DefaultLiteral: defaultLiteral,
		Kind:           KindPgtype,
		Supported:      true,
	}
}
