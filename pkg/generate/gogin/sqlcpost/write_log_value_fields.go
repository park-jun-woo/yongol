//ff:func feature=gen-gogin type=generator control=iteration dimension=1
//ff:what writeLogValueFields — LogValue() 본문의 slog.GroupValue 필드 라인 기록

package sqlcpost

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// writeLogValueFields writes one slog attribute line per column. Sensitive
// columns are masked as slog.String(name, "[REDACTED]"); other columns use
// the slog constructor matching their Go type (see slogAttrLine).
func writeLogValueFields(b *strings.Builder, t ddl.Table, cols []string) {
	for _, col := range cols {
		goType := t.Columns[col]
		fieldName := sqlcFieldName(col)
		if t.SensitiveColumns[col] {
			b.WriteString(fmt.Sprintf("\t\tslog.String(%q, \"[REDACTED]\"),\n", col))
			continue
		}
		b.WriteString("\t\t" + slogAttrLine(col, fieldName, goType) + ",\n")
	}
}
