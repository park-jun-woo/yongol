//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what findDDLTableByModelName — sqlc 모델명 → DDL 테이블 조회

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// findDDLTableByModelName looks up a DDL table by its sqlc model name. The
// parser stores the raw table name ("workflows"); the convert emitter
// receives the PascalCase sqlc model name ("Workflow"). We normalise both
// sides to the same lower-snake singular form ("workflow") for comparison.
func findDDLTableByModelName(tables []ddl.Table, modelName string) *ddl.Table {
	target := ddlTableSingular(pascalToSnake(modelName))
	for i := range tables {
		t := &tables[i]
		lower := strings.ToLower(t.Name)
		if ddlTableSingular(lower) == target {
			return t
		}
	}
	return nil
}
