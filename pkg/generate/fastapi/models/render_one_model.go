//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what renderOneModel — 단일 DDL Table → SQLAlchemy model class 렌더링 (FK/Index/Unique 포함)

package models

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// renderOneModel writes a single SQLAlchemy model class for a DDL table.
// Includes ForeignKey in column definitions and Index/UniqueConstraint in
// __table_args__.
func renderOneModel(b *strings.Builder, table ddl.Table) error {
	className := pascalCase(table.Name)
	b.WriteString(fmt.Sprintf("class %s(Base):\n", className))
	b.WriteString(fmt.Sprintf("    __tablename__ = \"%s\"\n\n", table.Name))

	for _, colName := range table.ColumnOrder {
		col, ok := table.Columns[colName]
		if !ok {
			continue
		}
		renderColumn(b, col, colName, table.PrimaryKey, table.ForeignKeys)
	}

	// __table_args__ for indexes and unique constraints.
	args := renderTableArgs(table)
	if args != "" {
		b.WriteString("\n")
		b.WriteString(args)
	}

	return nil
}
