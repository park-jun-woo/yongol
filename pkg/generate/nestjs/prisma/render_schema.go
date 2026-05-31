//ff:func feature=gen-nestjs type=generator control=iteration dimension=1
//ff:what RenderSchema — DDL Table 배열 → Prisma schema.prisma 소스 생성 (역관계 포함)

package prisma

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// RenderSchema produces a Prisma schema file content from DDL tables.
// It generates the datasource, generator, model blocks, and reverse
// relations for FK targets.
func RenderSchema(tables []ddl.Table) (string, error) {
	// Build reverse relation map: refTable → []reverseRelation.
	reverseMap := buildReverseRelations(tables)

	var b strings.Builder
	writeSchemaHeader(&b)

	for i, table := range tables {
		revRels := reverseMap[table.Name]
		if err := renderModel(&b, table, revRels); err != nil {
			return "", fmt.Errorf("RenderSchema(%s): %w", table.Name, err)
		}
		if i < len(tables)-1 {
			b.WriteString("\n")
		}
	}

	return b.String(), nil
}
