//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what renderTableArgs — DDL Index/UniqueConstraint → __table_args__ 튜플 생성

package models

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// renderTableArgs builds the __table_args__ tuple for Index and UniqueConstraint.
func renderTableArgs(table ddl.Table) string {
	var entries []string
	for _, idx := range table.Indexes {
		cols := formatColumnList(idx.Columns)
		if idx.IsUnique {
			entries = append(entries, fmt.Sprintf("        UniqueConstraint(%s, name=\"%s\")", cols, idx.Name))
		} else {
			entries = append(entries, fmt.Sprintf("        Index(\"%s\", %s)", idx.Name, cols))
		}
	}
	if len(entries) == 0 {
		return ""
	}
	var b strings.Builder
	b.WriteString("    __table_args__ = (\n")
	for _, e := range entries {
		b.WriteString(e + ",\n")
	}
	b.WriteString("    )\n")
	return b.String()
}
