//ff:func feature=gen-gogin type=util control=selection
//ff:what pgtypeTimestamp — TIMESTAMPTZ/TIMESTAMP/DATE 컬럼의 pgtype.Timestamptz 매핑

package types

// pgtypeTimestamp returns the binding for a temporal column. sqlc pgx/v5
// emits pgtype.Timestamptz / pgtype.Timestamp / pgtype.Date depending on
// the source PG type; for the convert / response sites we uniformly read
// `.Time` because all three wrappers expose the value via that field.
//
// NeedsOverride is true because the default sqlc map for `timestamp /
// timestamptz / date` predates pgx/v5 in some configurations and the
// project standardises on pgtype.Timestamptz to round-trip TZ correctly.
func pgtypeTimestamp(head string, notNull bool, defaultLiteral string) GoTypeBinding {
	sqlcType := "pgtype.Timestamptz"
	switch head {
	case "TIMESTAMP":
		sqlcType = "pgtype.Timestamp"
	case "DATE":
		sqlcType = "pgtype.Date"
	}
	apiField := "time.Time"
	if !notNull {
		apiField = "*time.Time"
	}
	return GoTypeBinding{
		SqlcGoType:    sqlcType,
		NeedsOverride: true,
		ApiField:      apiField,
		Imports: []string{
			"github.com/jackc/pgx/v5/pgtype",
			"time",
		},
		ConvertExpr:    "{row}.{field}.Time",
		InsertExpr:     "{var}",
		ResponseExpr:   "{var}.{field}.Time",
		DefaultLiteral: defaultLiteral,
		Kind:           KindPgtype,
		Supported:      true,
	}
}
