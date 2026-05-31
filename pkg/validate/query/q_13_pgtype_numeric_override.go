//ff:func feature=validate type=rule control=sequence topic=query-structural
//ff:what Q-13 — DDL 에 NUMERIC/DECIMAL 컬럼이 있으면 sqlc.yaml 에 pgtype.Numeric overrides 강제

package query

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// q13PgtypeNumericOverride validates Q-13: any DDL NUMERIC / DECIMAL
// column requires both nullable=false and nullable=true pgtype.Numeric
// overrides in sqlc.yaml. Without the override sqlc emits raw
// pgtype.Numeric in some configurations and a parsed `interface{}` in
// others, breaking convert sites.
func q13PgtypeNumericOverride(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	return checkPgtypeOverride(fs, pgtypeOverrideRule{
		RuleID:    "Q-13",
		DBType:    "numeric",
		PgPackage: "pgtype",
		PgType:    "Numeric",
		Filter:    isNumericColumn,
		Advice: "Add to sql[].gen.go.overrides:\n" +
			"  - db_type: \"numeric\"\n" +
			"    nullable: false\n" +
			"    go_type:\n" +
			"      import: \"github.com/jackc/pgx/v5/pgtype\"\n" +
			"      type: \"Numeric\"\n" +
			"  - db_type: \"numeric\"\n" +
			"    nullable: true\n" +
			"    go_type:\n" +
			"      import: \"github.com/jackc/pgx/v5/pgtype\"\n" +
			"      type: \"Numeric\"\n" +
			"NUMERIC / DECIMAL have no Go native type — yongol routes them through pgNumericToString in the convert layer.",
	})
}
