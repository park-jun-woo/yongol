//ff:func feature=agent type=helper control=sequence
//ff:what buildDDLUserPrompt — DDL 생성용 user prompt 구성

package agent

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

// buildDDLUserPrompt builds the user prompt for generating a DDL file.
func buildDDLUserPrompt(tableName string, td features.TableDef, feats []features.Feature) string {
	var b strings.Builder

	fmt.Fprintf(&b, "Table: %s\n", tableName)

	if len(td.BelongsTo) > 0 {
		fmt.Fprintf(&b, "belongs_to: %s\n", strings.Join(td.BelongsTo, ", "))
	}
	if len(td.HasMany) > 0 {
		fmt.Fprintf(&b, "has_many: %s\n", strings.Join(td.HasMany, ", "))
	}
	if len(td.States) > 0 {
		fmt.Fprintf(&b, "states: %s\n", strings.Join(td.States, ", "))
	}

	if len(feats) > 0 {
		b.WriteString("\nRelated features:\n")
		for _, f := range feats {
			fmt.Fprintf(&b, "  - %s %s: %s\n", f.Op, f.Path, f.Desc)
		}
	}

	b.WriteString("\nGenerate a PostgreSQL CREATE TABLE statement for this table.")
	b.WriteString("\nOutput ONLY the SQL. No explanations. No markdown fences.")

	return b.String()
}
