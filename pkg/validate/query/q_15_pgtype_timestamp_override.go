//ff:func feature=validate type=rule control=sequence topic=query-structural
//ff:what Q-15 — DDL 에 TIMESTAMP (no TZ) 컬럼이 있으면 sqlc.yaml 에 pgtype.Timestamp overrides 강제

package query

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// q15PgtypeTimestampOverride validates Q-15: any DDL TIMESTAMP (without
// time zone) column requires both pgtype.Timestamp overrides. Distinct
// from Q-14 because the underlying pgtype is different.
func q15PgtypeTimestampOverride(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	return checkPgtypeOverride(fs, pgtypeOverrideRule{
		RuleID:    "Q-15",
		DBType:    "timestamp",
		PgPackage: "pgtype",
		PgType:    "Timestamp",
		Filter:    isTimestampColumn,
		Advice: "Add to sql[].gen.go.overrides:\n" +
			"  - db_type: \"timestamp\"\n" +
			"    nullable: false\n" +
			"    go_type:\n" +
			"      import: \"github.com/jackc/pgx/v5/pgtype\"\n" +
			"      package: \"pgtype\"\n" +
			"      type: \"Timestamp\"\n" +
			"  - db_type: \"timestamp\"\n" +
			"    nullable: true\n" +
			"    go_type:\n" +
			"      import: \"github.com/jackc/pgx/v5/pgtype\"\n" +
			"      package: \"pgtype\"\n" +
			"      type: \"Timestamp\"\n",
	})
}

