//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what renderModel — 단일 DDL Table → Prisma model 블록 렌더링

package prisma

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// renderModel writes a single Prisma model block for a DDL table.
func renderModel(b *strings.Builder, table ddl.Table) error {
	modelName := pascalCase(table.Name)
	b.WriteString(fmt.Sprintf("model %s {\n", modelName))

	for _, colName := range table.ColumnOrder {
		col, ok := table.Columns[colName]
		if !ok {
			continue
		}
		renderColumn(b, col, colName, table.PrimaryKey)
	}

	b.WriteString("}\n")
	return nil
}
