//ff:func feature=gen-ir type=util control=iteration dimension=1
//ff:what matchDDLColumns -- FieldArg.Key 를 DDL 테이블 컬럼과 매칭해 ColumnName/IsPK 세팅

package ir

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/util/caseconv"
)

// matchDDLColumns sets ColumnName and IsPK on each FieldArg in args whose
// snake_case Key corresponds to a column in tbl.
func matchDDLColumns(args []FieldArg, tbl *ddl.Table, pkSet map[string]bool) {
	for j := range args {
		snake := caseconv.PascalToSnake(args[j].Key)
		if _, ok := tbl.Columns[snake]; ok {
			args[j].ColumnName = snake
			args[j].IsPK = pkSet[snake]
		}
	}
}
