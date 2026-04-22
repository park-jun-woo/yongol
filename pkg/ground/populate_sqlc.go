//ff:func feature=rule type=loader control=iteration dimension=1
//ff:what populateSQLc — DDL 컬럼명을 모델별 sqlc 파라미터 집합으로 등록 + sqlc row type 노출
package ground

import (
	"github.com/jinzhu/inflection"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/util/caseconv"
)

// populateSQLc registers column names as sqlc parameter set per model.
// XQS-14 compares SSaC input key case against this set. Columns are stored
// snake_case; the consumer normalizes both sides for comparison.
//
// Additionally, each non-empty QuerySpec.RowType is exposed as a singleton
// Lookup entry under "SQLc.rowType.<RowType>". Downstream rules (e.g. XDS-12)
// can detect sqlc-synthesized row structs without walking fs.SQLcQueries.
func populateSQLc(g *rule.Ground, fs *yongol.Fullstack) {
	for _, t := range fs.DDLTables {
		model := caseconv.SnakeToPascal(inflection.Singular(t.Name))
		params := make(rule.StringSet, len(t.Columns))
		for col := range t.Columns {
			params[col] = true
		}
		g.Lookup["SQLc.param."+model] = params
	}
	for _, q := range fs.SQLcQueries {
		if q.RowType == "" {
			continue
		}
		g.Lookup["SQLc.rowType."+q.RowType] = rule.StringSet{q.RowType: true}
	}
}
