//ff:func feature=gen-nestjs type=util control=sequence
//ff:what renderColumn — 단일 DDL Column → Prisma 필드 한 줄 렌더링

package prisma

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// renderColumn writes a single Prisma field line for a DDL column.
func renderColumn(b *strings.Builder, col ddl.Column, colName string, primaryKey []string, indexes ...[]ddl.Index) {
	prismaType := pgToPrismaType(col.RawType)
	optional := ""
	if !col.NotNull && !isPrimaryKey(colName, primaryKey) {
		optional = "?"
	}
	attrs := columnAttributes(col, colName, primaryKey)
	// Check single-column unique indexes.
	if len(indexes) > 0 && isColumnUnique(colName, indexes[0]) {
		if attrs != "" {
			attrs += " @unique"
		} else {
			attrs = "@unique"
		}
	}
	if attrs != "" {
		attrs = " " + attrs
	}
	b.WriteString(fmt.Sprintf("  %-20s %s%s%s\n", colName, prismaType, optional, attrs))
}
