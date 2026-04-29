//ff:func feature=validate type=rule control=sequence topic=query-structural
//ff:what Q-17 — DDL 에 INET/CIDR 컬럼이 있으면 sqlc.yaml 에 pgtype.Inet overrides 강제

package query

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// q17PgtypeInetOverride validates Q-17: any DDL INET / CIDR column
// requires both pgtype.Inet overrides. The override is needed because
// sqlc pgx/v5 has no default mapping for `inet` / `cidr`.
func q17PgtypeInetOverride(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	return checkPgtypeOverride(fs, pgtypeOverrideRule{
		RuleID:    "Q-17",
		DBType:    "inet",
		PgPackage: "pgtype",
		PgType:    "Inet",
		Filter:    isInetColumn,
		Advice: "Add to sql[].gen.go.overrides:\n" +
			"  - db_type: \"inet\"\n" +
			"    nullable: false\n" +
			"    go_type:\n" +
			"      import: \"github.com/jackc/pgx/v5/pgtype\"\n" +
			"      package: \"pgtype\"\n" +
			"      type: \"Inet\"\n" +
			"  - db_type: \"inet\"\n" +
			"    nullable: true\n" +
			"    go_type:\n" +
			"      import: \"github.com/jackc/pgx/v5/pgtype\"\n" +
			"      package: \"pgtype\"\n" +
			"      type: \"Inet\"\n" +
			"yongol routes INET / CIDR through pgInetToString in the convert layer.",
	})
}

