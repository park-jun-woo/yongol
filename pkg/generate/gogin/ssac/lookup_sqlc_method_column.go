//ff:func feature=gen-gogin type=util control=sequence
//ff:what lookupSQLCMethodColumn — sqlc method 의 model + 컬럼명 → DDL Column (없으면 nil)

package ssac

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// lookupSQLCMethodColumn returns the DDL Column referenced by the sqlc
// params field name for the currently-active method. nil when:
//
//   - g.activeMethod is empty (no method context was set)
//   - the sqlc query catalogue has no entry with the method name
//   - the table cannot be resolved from the model name
//   - the column does not exist
//
// activeMethod is set by sqlcArgs / sqlcArgsSingle / sqlcArgsMulti so
// per-column lookups can resolve the target DDL table from
// SQLcQueries.
func (g *methodGen) lookupSQLCMethodColumn(paramKey string) *ddl.Column {
	if g.activeMethod == "" {
		return nil
	}
	model := g.modelForSQLCMethod(g.activeMethod)
	if model == "" {
		return nil
	}
	return lookupDDLColumn(g.DDLTables, model, paramKey)
}
