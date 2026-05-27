//ff:func feature=gen-ir type=util control=iteration dimension=1
//ff:what findDDLTable -- sqlc 모델명 → DDL 테이블 조회 (gogin/ssac/findDDLTableByModelName 미러)

package ir

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/util/caseconv"
)

// findDDLTable looks up a DDL table by its sqlc model name using the same
// singular-matching logic as gogin's findDDLTableByModelName.
func findDDLTable(tables []ddl.Table, modelName string) *ddl.Table {
	target := ddlTableSingularIR(caseconv.PascalToSnake(modelName))
	for i := range tables {
		t := &tables[i]
		lower := strings.ToLower(t.Name)
		if ddlTableSingularIR(lower) == target {
			return t
		}
	}
	return nil
}
