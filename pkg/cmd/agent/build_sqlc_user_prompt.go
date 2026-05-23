//ff:func feature=agent type=helper control=sequence
//ff:what buildSQLcUserPrompt — sqlc 쿼리 생성용 user prompt 구성

package agent

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

// buildSQLcUserPrompt builds the user prompt for generating sqlc queries.
func buildSQLcUserPrompt(tableName, ddlContent string, feats []features.Feature) string {
	var b strings.Builder

	fmt.Fprintf(&b, "Table: %s\n\n", tableName)
	fmt.Fprintf(&b, "DDL:\n%s\n\n", ddlContent)

	if len(feats) > 0 {
		b.WriteString("Related features with cardinality hints:\n")
		for _, f := range feats {
			hint := cardinalityHint(f.Op, f.Path)
			fmt.Fprintf(&b, "  - %s %s: %s (cardinality: %s)\n", f.Op, f.Path, f.Desc, hint)
		}
		b.WriteByte('\n')
	}

	b.WriteString("Generate sqlc-compatible SQL queries for this table.")
	b.WriteString("\nOutput ONLY the SQL. No explanations. No markdown fences.")

	return b.String()
}
