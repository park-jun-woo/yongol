//ff:func feature=agent type=helper control=sequence
//ff:what buildStateMachineUserPrompt — state diagram 생성용 user prompt 구성

package agent

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

// buildStateMachineUserPrompt builds the user prompt for generating a state diagram.
func buildStateMachineUserPrompt(tableName string, states []string, feats []features.Feature) string {
	var b strings.Builder

	fmt.Fprintf(&b, "Table: %s\n", tableName)
	fmt.Fprintf(&b, "States: %s\n", strings.Join(states, ", "))

	if len(feats) > 0 {
		b.WriteString("\nRelated features:\n")
		for _, f := range feats {
			fmt.Fprintf(&b, "  - %s %s: %s\n", f.Op, f.Path, f.Desc)
		}
	}

	b.WriteString("\nGenerate a Mermaid stateDiagram-v2 for this table.")
	b.WriteString("\nOutput a markdown file with '# <table_name>' heading followed by a mermaid fenced code block.")
	b.WriteString("\nInclude [*] --> first_state as the initial transition.")
	b.WriteString("\nUse operationId names for transition labels where applicable.")

	return b.String()
}
