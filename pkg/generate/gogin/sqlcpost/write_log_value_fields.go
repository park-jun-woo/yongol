//ff:func feature=gen-gogin type=generator control=iteration dimension=1
//ff:what writeLogValueFields — LogValue() 본문의 slog.GroupValue 필드 라인 기록

package sqlcpost

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// writeLogValueFields writes one slog attribute line per column. Sensitive
// columns are masked as slog.String(name, "[REDACTED]"); all other columns
// go through slogAttrLine which emits slog.Any(name, r.Field) uniformly
// (see slogAttrLine for rationale / BUG-024).
func writeLogValueFields(b *strings.Builder, t ddl.Table, cols []string) {
	for _, col := range cols {
		c := t.Columns[col]
		goType := ddl.GoTypeOf(c)
		fieldName := sqlcFieldName(col)
		if c.Sensitive {
			b.WriteString(fmt.Sprintf("\t\tslog.String(%q, \"[REDACTED]\"),\n", col))
			continue
		}
		b.WriteString("\t\t" + slogAttrLine(col, fieldName, goType) + ",\n")
	}
}
