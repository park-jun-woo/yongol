//ff:func feature=gen-gogin type=util control=sequence
//ff:what lookupPKColumn — target 변수의 모델명으로 PK(id) DDL 컬럼 조회

package ssac

import "github.com/park-jun-woo/yongol/pkg/parser/ddl"

// lookupPKColumn resolves the primary key (id) column for the model that
// the target variable refers to. Uses VarTypes to map variable → model
// name, then looks up the DDL table. Returns nil when the model cannot be
// resolved or the table has no "id" column.
func (g *methodGen) lookupPKColumn(target string) *ddl.Column {
	modelName := g.VarTypes[target]
	if modelName == "" {
		return nil
	}
	return lookupDDLColumn(g.DDLTables, modelName, "id")
}
