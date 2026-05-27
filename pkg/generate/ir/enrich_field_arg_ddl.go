//ff:func feature=gen-ir type=util control=iteration dimension=2
//ff:what enrichFieldArgDDL -- Pass 1: 모든 Op 의 FieldArg.SourceColumn 세팅 / Pass 2: CRUD Op DDL 테이블 매칭으로 ColumnName/IsPK 세팅

package ir

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/util/caseconv"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// enrichFieldArgDDL runs two passes over all ops:
//
// Pass 1 sets SourceColumn on every FieldArg that has a Field accessor,
// regardless of Op kind. This ensures @call/@eval/@auth ops get proper
// snake_case source column names even without DDL table matching.
//
// Pass 2 enriches CRUD ops (get/post/put/delete) with DDL table metadata,
// setting ColumnName and IsPK by matching the argument Key against DDL
// table columns.
func enrichFieldArgDDL(ops []Op, fs *yongol.Fullstack) {
	// Pass 1: SourceColumn for ALL ops (DDL-independent).
	for i := range ops {
		for _, args := range collectFieldArgSlices(&ops[i]) {
			for j := range *args {
				fa := &(*args)[j]
				if fa.Field != "" {
					fa.SourceColumn = caseconv.PascalToSnake(
						strings.TrimPrefix(fa.Field, "."))
				}
			}
		}
	}

	// Pass 2: DDL table matching for CRUD ops (ColumnName/IsPK).
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
