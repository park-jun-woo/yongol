//ff:func feature=gen-ir type=util control=iteration dimension=1
//ff:what enrichOpDDLColumns -- 단일 CRUD Op 을 DDL 테이블 매칭으로 ColumnName/IsPK 세팅

package ir

import "github.com/park-jun-woo/yongol/pkg/yongol"

// enrichOpDDLColumns matches op's model against a DDL table and sets
// ColumnName/IsPK on each FieldArg whose Key maps to a table column (Pass 2).
func enrichOpDDLColumns(op *Op, fs *yongol.Fullstack) {
	modelName := opModelName(op)
	if modelName == "" {
		return
	}
	tbl := findDDLTable(fs.DDLTables, modelName)
	if tbl == nil {
		return
	}
	pkSet := pkColumnSet(tbl.PrimaryKey)
	for _, args := range collectFieldArgSlices(op) {
		matchDDLColumns(*args, tbl, pkSet)
	}
}
