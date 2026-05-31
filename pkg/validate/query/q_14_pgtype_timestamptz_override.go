//ff:func feature=validate type=rule control=sequence topic=query-structural
//ff:what Q-14 — DDL 에 TIMESTAMPTZ 컬럼이 있으면 sqlc.yaml 에 pgtype.Timestamptz overrides 강제

package query

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// q14PgtypeTimestamptzOverride validates Q-14: any DDL TIMESTAMPTZ
// column requires both nullable=false and nullable=true pgtype.Timestamptz
// overrides in sqlc.yaml. Without the explicit override sqlc may pick
// time.Time which loses TZ semantics on round trip.
func q14PgtypeTimestamptzOverride(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	return checkPgtypeOverride(fs, pgtypeOverrideRule{
		RuleID:    "Q-14",
		DBType:    "timestamptz",
		PgPackage: "pgtype",
		PgType:    "Timestamptz",
		Filter:    isTimestamptzColumn,
		Advice: "Add to sql[].gen.go.overrides:\n" +
			"  - db_type: \"timestamptz\"\n" +
			"    nullable: false\n" +
			"    go_type:\n" +
			"      import: \"github.com/jackc/pgx/v5/pgtype\"\n" +
			"      type: \"Timestamptz\"\n" +
			"  - db_type: \"timestamptz\"\n" +
			"    nullable: true\n" +
			"    go_type:\n" +
			"      import: \"github.com/jackc/pgx/v5/pgtype\"\n" +
			"      type: \"Timestamptz\"\n",
	})
}
