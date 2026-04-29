//ff:func feature=validate type=rule control=sequence topic=query-structural
//ff:what Q-18 — DDL 에 INTERVAL 컬럼이 있으면 sqlc.yaml 에 pgtype.Interval overrides 강제

package query

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// q18PgtypeIntervalOverride validates Q-18: any DDL INTERVAL column
// requires both pgtype.Interval overrides. sqlc pgx/v5 has no default
// mapping; pgtype.Interval exposes Microseconds + Days + Months which
// the convert layer stringifies via pgIntervalToString.
func q18PgtypeIntervalOverride(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	return checkPgtypeOverride(fs, pgtypeOverrideRule{
		RuleID:    "Q-18",
		DBType:    "interval",
		PgPackage: "pgtype",
		PgType:    "Interval",
		Filter:    isIntervalColumn,
		Advice: "Add to sql[].gen.go.overrides:\n" +
			"  - db_type: \"interval\"\n" +
			"    nullable: false\n" +
			"    go_type:\n" +
			"      import: \"github.com/jackc/pgx/v5/pgtype\"\n" +
			"      package: \"pgtype\"\n" +
			"      type: \"Interval\"\n" +
			"  - db_type: \"interval\"\n" +
			"    nullable: true\n" +
			"    go_type:\n" +
			"      import: \"github.com/jackc/pgx/v5/pgtype\"\n" +
			"      package: \"pgtype\"\n" +
			"      type: \"Interval\"\n",
	})
}

