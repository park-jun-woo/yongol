//ff:func feature=rule type=loader control=iteration dimension=1
//ff:what populateSymbolMethods — 모델별 표준 CRUD + sqlc 쿼리 메서드 집합을 Ground에 등록
package ground

import (
	"github.com/jinzhu/inflection"

	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/util/caseconv"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// populateSymbolMethods registers method names per model from two sources:
//  1. Default CRUD names (Get/List/Create/...) that conventional repository
//     generators emit by name.
//  2. Parsed sqlc QuerySpecs carried in fs.SQLcQueries. Each spec contributes
//     its Method (already stripped of the model prefix) to the model's set.
//
// S-49 (SSaC @call Model.Method existence) consumes this set.
func populateSymbolMethods(g *rule.Ground, fs *yongol.Fullstack) {
	defaults := []string{
		"Get", "List", "Create", "Update", "Delete", "Exists", "Count",
		"Find", "FindByID", "FindOne", "Insert", "Remove",
	}
	for _, t := range fs.DDLTables {
		model := caseconv.SnakeToPascal(inflection.Singular(t.Name))
		methods := make(rule.StringSet, len(defaults))
		for _, m := range defaults {
			methods[m] = true
		}
		g.Lookup["SymbolTable.method."+model] = methods
	}
	for _, q := range fs.SQLcQueries {
		if q.Model == "" || q.Method == "" {
			continue
		}
		key := "SymbolTable.method." + q.Model
		methods := g.Lookup[key]
		if methods == nil {
			methods = make(rule.StringSet)
			g.Lookup[key] = methods
		}
		methods[q.Method] = true
	}
}
