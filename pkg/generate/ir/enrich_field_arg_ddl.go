//ff:func feature=gen-ir type=util control=iteration dimension=2
//ff:what enrichFieldArgDDL -- DDL 테이블 컬럼 매칭으로 모든 FieldArg 의 ColumnName/IsPK 세팅

package ir

import (
	"github.com/park-jun-woo/yongol/pkg/util/caseconv"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// enrichFieldArgDDL populates ColumnName and IsPK on every FieldArg by
// looking up the argument's Key against the DDL table columns. The table
// is resolved from the owning Op's Model field via the same singular
// matching logic used by gogin's findDDLTableByModelName.
func enrichFieldArgDDL(ops []Op, fs *yongol.Fullstack) {
	if fs == nil || len(fs.DDLTables) == 0 {
		return
	}
	for i := range ops {
		modelName := opModelName(&ops[i])
		if modelName == "" {
			continue
		}
		tbl := findDDLTable(fs.DDLTables, modelName)
		if tbl == nil {
			continue
		}
		pkSet := make(map[string]bool, len(tbl.PrimaryKey))
		for _, pk := range tbl.PrimaryKey {
			pkSet[pk] = true
		}
		for _, args := range collectFieldArgSlices(&ops[i]) {
			for j := range *args {
				fa := &(*args)[j]
				snake := caseconv.PascalToSnake(fa.Key)
				if _, ok := tbl.Columns[snake]; ok {
					fa.ColumnName = snake
					fa.IsPK = pkSet[snake]
				}
			}
		}
	}
}
