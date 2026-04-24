//ff:func feature=validate type=rule control=sequence topic=query-structural
//ff:what diagnoseSqlcPackageEntry — sqlc.yaml sql[i] 의 sql_package 값에 대한 Q-11 진단 생성

package query

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// diagnoseSqlcPackageEntry returns a Q-11 diagnostic for a single
// sql[i].gen.go.sql_package value, or nil when the value is the required
// "pgx/v5". Extracted from q11SqlPackagePgxV5 to keep that func's range
// body under the Q4 PURE line budget.
func diagnoseSqlcPackageEntry(i int, pkg string) *diagnostic.Diagnostic {
	if pkg == "pgx/v5" {
		return nil
	}
	current := pkg
	if current == "" {
		current = "(absent; sqlc defaults to database/sql)"
	} else {
		current = fmt.Sprintf("%q", current)
	}
	return &diagnostic.Diagnostic{
		File:  "db/sqlc.yaml",
		Line:  0,
		Phase: diagnostic.PhaseValidate,
		Level: diagnostic.LevelError,
		Message: fmt.Sprintf(
			"[Q-11] sqlc.yaml sql[%d].gen.go.sql_package must be \"pgx/v5\" (current: %s)",
			i, current,
		),
		Advice: "yongol's backend codegen is unified on pgx/v5. Update db/sqlc.yaml:\n" +
			"  gen:\n" +
			"    go:\n" +
			"      sql_package: pgx/v5\n" +
			"Then re-run `yongol generate <specs> <arts>`. database/sql / pgx/v4 support was removed.",
	}
}
