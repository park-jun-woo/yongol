//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what renderModel — 단일 DDL Table → Prisma model 블록 렌더링 (역관계 + @@map 일괄)

package prisma

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// renderModel writes a single Prisma model block for a DDL table.
// Model name is singular PascalCase with @@map to the original table name
// (always applied). FK relations, reverse relations, and indexes are
// rendered after columns.
func renderModel(b *strings.Builder, table ddl.Table, reverseRels []reverseRelation) error {
	singularName := singularize(table.Name)
	modelName := pascalCase(singularName)
	b.WriteString(fmt.Sprintf("model %s {\n", modelName))

	for _, colName := range table.ColumnOrder {
		col, ok := table.Columns[colName]
		if !ok {
			continue
		}
		renderColumn(b, col, colName, table.PrimaryKey, table.Indexes)
	}

	// Render @relation fields for foreign keys (forward relations).
	for _, fk := range table.ForeignKeys {
		refModel := pascalCase(singularize(fk.RefTable))
		relField := strings.TrimSuffix(fk.Column, "_id")
		if relField == fk.Column {
			relField = strings.TrimSuffix(fk.Column, "_ID")
		}
		b.WriteString(fmt.Sprintf("  %-20s %s @relation(fields: [%s], references: [%s])\n",
			relField, refModel, fk.Column, fk.RefColumn))
	}

	// Render reverse (has-many) relations for models referenced by other tables.
	for _, rev := range dedupReverseRelations(reverseRels) {
		b.WriteString(fmt.Sprintf("  %-20s %s[]\n", rev.FieldName, rev.ModelName))
	}

	// Render @@index for non-unique indexes.
	for _, idx := range table.Indexes {
		if idx.IsUnique {
			continue
		}
		cols := strings.Join(idx.Columns, ", ")
		b.WriteString(fmt.Sprintf("  @@index([%s])\n", cols))
	}

	// Render @@unique for composite unique indexes (single-column unique is
	// handled by @unique on the column itself).
	for _, idx := range table.Indexes {
		if !idx.IsUnique || len(idx.Columns) <= 1 {
			continue
		}
		cols := strings.Join(idx.Columns, ", ")
		b.WriteString(fmt.Sprintf("  @@unique([%s])\n", cols))
	}

	// Always map to the original DDL table name.
	b.WriteString(fmt.Sprintf("  @@map(\"%s\")\n", table.Name))

	b.WriteString("}\n")
	return nil
}
