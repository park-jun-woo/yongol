//ff:func feature=validate type=util control=sequence topic=ssac-structural
//ff:what xss60InferDotExpr — dot expression (var.Field) 에서 DDL 컬럼 타입을 역추적하여 Go 타입 반환

package ssac

import (
	"github.com/ettle/strcase"

	"github.com/park-jun-woo/yongol/pkg/generate/gogin/types"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// xss60InferDotExpr resolves a dot expression (e.g. "user.Email") to a Go type
// by looking up the variable's model type, then the DDL table column.
func xss60InferDotExpr(expr string, dotIdx int, fn parsessac.ServiceFunc, tableMap map[string]*ddl.Table) string {
	varName := expr[:dotIdx]
	colName := expr[dotIdx+1:]

	modelName := xss60ResolveVarModel(varName, fn)
	if modelName == "" {
		return ""
	}

	tableName := xss60ModelToTableName(modelName)
	table, ok := tableMap[tableName]
	if !ok {
		return ""
	}
	// DDL column names are snake_case; SSaC field names are PascalCase.
	ddlColName := strcase.ToSnake(colName)
	col, ok := table.Columns[ddlColName]
	if !ok {
		return ""
	}
	return types.GoTypeOf(col)
}
