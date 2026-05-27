//ff:func feature=gen-nestjs type=generator control=iteration dimension=1
//ff:what RenderSchema — DDL Table 배열 → Prisma schema.prisma 소스 생성 (역관계 포함)

package prisma

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// reverseRelation tracks a reverse (has-many) relation that a referenced
// model needs for Prisma completeness.
type reverseRelation struct {
	FieldName string // plural snake_case (e.g. "posts")
	ModelName string // PascalCase referring model (e.g. "Post")
}

// RenderSchema produces a Prisma schema file content from DDL tables.
// It generates the datasource, generator, model blocks, and reverse
// relations for FK targets.
func RenderSchema(tables []ddl.Table) (string, error) {
	// Build reverse relation map: refTable → []reverseRelation.
	reverseMap := buildReverseRelations(tables)

	var b strings.Builder

	b.WriteString("datasource db {\n")
	b.WriteString("  provider = \"postgresql\"\n")
	b.WriteString("  url      = env(\"DATABASE_URL\")\n")
	b.WriteString("}\n\n")

	b.WriteString("generator client {\n")
	b.WriteString("  provider = \"prisma-client-js\"\n")
	b.WriteString("}\n\n")

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

// buildReverseRelations scans all tables' ForeignKeys and returns a map
// from referenced table name to the list of reverse relation entries.
func buildReverseRelations(tables []ddl.Table) map[string][]reverseRelation {
	rm := make(map[string][]reverseRelation)
	for _, table := range tables {
		for _, fk := range table.ForeignKeys {
			refModelName := pascalCase(singularize(fk.RefTable))
			sourceModelName := pascalCase(singularize(table.Name))
			// Reverse field name: pluralized source table name.
			fieldName := table.Name
			rm[fk.RefTable] = append(rm[fk.RefTable], reverseRelation{
				FieldName: fieldName,
				ModelName: sourceModelName,
			})
			_ = refModelName // used implicitly via fk.RefTable
		}
	}
	return rm
}

// uniqueReverseFieldName returns a deduplicated field name when multiple
// FKs point to the same target. Currently uses table name as-is since
// each table has a unique name.
func uniqueReverseFieldName(base string, _ int) string {
	return base
}

// dedupReverseRelations removes duplicate entries (same source model).
func dedupReverseRelations(rels []reverseRelation) []reverseRelation {
	seen := make(map[string]bool)
	var result []reverseRelation
	for _, r := range rels {
		key := r.FieldName + ":" + r.ModelName
		if seen[key] {
			continue
		}
		seen[key] = true
		result = append(result, r)
	}
	return result
}

var _ = strings.HasSuffix // keep import
