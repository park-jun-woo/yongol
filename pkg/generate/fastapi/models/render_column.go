//ff:func feature=gen-fastapi type=util control=sequence
//ff:what renderColumn — 단일 DDL Column → SQLAlchemy mapped_column 한 줄 렌더링

package models

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// renderColumn writes a single mapped_column line for a DDL column.
func renderColumn(b *strings.Builder, col ddl.Column, colName string, primaryKey []string) {
	saType := mapPGToSA(col.RawType)
	pyType := mapPGToPython(col.RawType, col.NotNull)

	var attrs []string
	attrs = append(attrs, saType)

	if isPrimaryKey(colName, primaryKey) {
		attrs = append(attrs, "primary_key=True")
	}
	if !col.NotNull && !isPrimaryKey(colName, primaryKey) {
		attrs = append(attrs, "nullable=True")
	}
	if col.HasDefault {
		def := saDefault(col)
		if def != "" {
			attrs = append(attrs, def)
		}
	}

	b.WriteString(fmt.Sprintf("    %s: Mapped[%s] = mapped_column(%s)\n",
		colName, pyType, strings.Join(attrs, ", ")))
}
