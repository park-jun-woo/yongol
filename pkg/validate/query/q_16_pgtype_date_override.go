//ff:func feature=validate type=rule control=sequence topic=query-structural
//ff:what Q-16 — DDL 에 DATE 컬럼이 있으면 sqlc.yaml 에 pgtype.Date overrides 강제

package query

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// q16PgtypeDateOverride validates Q-16: any DDL DATE column requires
// both pgtype.Date overrides. Distinct from Q-14 / Q-15 because the
// underlying pgtype is different (no time component, no TZ).
func q16PgtypeDateOverride(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	return checkPgtypeOverride(fs, pgtypeOverrideRule{
		RuleID:    "Q-16",
		DBType:    "date",
		PgPackage: "pgtype",
		PgType:    "Date",
		Filter:    isDateColumn,
		Advice: "Add to sql[].gen.go.overrides:\n" +
			"  - db_type: \"date\"\n" +
			"    nullable: false\n" +
			"    go_type:\n" +
			"      import: \"github.com/jackc/pgx/v5/pgtype\"\n" +
			"      package: \"pgtype\"\n" +
			"      type: \"Date\"\n" +
			"  - db_type: \"date\"\n" +
			"    nullable: true\n" +
			"    go_type:\n" +
			"      import: \"github.com/jackc/pgx/v5/pgtype\"\n" +
			"      package: \"pgtype\"\n" +
			"      type: \"Date\"\n",
	})
}

