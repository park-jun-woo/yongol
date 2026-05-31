//ff:func feature=gen-nestjs type=generator control=sequence
//ff:what writeSchemaHeader — Prisma datasource/generator 헤더 블록 출력

package prisma

import "strings"

// writeSchemaHeader writes the Prisma datasource and generator blocks.
func writeSchemaHeader(b *strings.Builder) {
	b.WriteString("datasource db {\n")
	b.WriteString("  provider = \"postgresql\"\n")
	b.WriteString("  url      = env(\"DATABASE_URL\")\n")
	b.WriteString("}\n\n")

	b.WriteString("generator client {\n")
	b.WriteString("  provider = \"prisma-client-js\"\n")
	b.WriteString("}\n\n")
}
