//ff:func feature=rule type=loader control=iteration dimension=1
//ff:what populateSymbolTable — DDL 테이블명→모델명 매핑 + sqlc 쿼리→메서드 매핑
package ground

import (
	"github.com/jinzhu/inflection"

	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/util/caseconv"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func populateSymbolTable(g *rule.Ground, fs *yongol.Fullstack) {
	models := make(rule.StringSet)
	for _, t := range fs.DDLTables {
		model := caseconv.SnakeToPascal(inflection.Singular(t.Name))
		models[model] = true
	}
	g.Lookup["SymbolTable.model"] = models
}
