//ff:func feature=gen-nestjs type=generator control=iteration dimension=1
//ff:what RenderSchema — DDL Table 배열 → Prisma schema.prisma 소스 생성

package prisma

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// RenderSchema produces a Prisma schema file content from DDL tables.
// It generates the datasource, generator, and model blocks.
func RenderSchema(tables []ddl.Table) (string, error) {
	var b strings.Builder

	b.WriteString("datasource db {\n")
	b.WriteString("  provider = \"postgresql\"\n")
	b.WriteString("  url      = env(\"DATABASE_URL\")\n")
	b.WriteString("}\n\n")

	b.WriteString("generator client {\n")
	b.WriteString("  provider = \"prisma-client-js\"\n")
	b.WriteString("}\n\n")

	for i, table := range tables {
		if err := renderModel(&b, table); err != nil {
			return "", fmt.Errorf("RenderSchema(%s): %w", table.Name, err)
		}
		if i < len(tables)-1 {
			b.WriteString("\n")
		}
	}

	return b.String(), nil
}
