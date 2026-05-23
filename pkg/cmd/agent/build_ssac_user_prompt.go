//ff:func feature=agent type=helper control=sequence
//ff:what buildSSaCUserPrompt — SSaC 서비스 생성용 user prompt 구성

package agent

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

// buildSSaCUserPrompt builds the user prompt for generating a single SSaC file.
func buildSSaCUserPrompt(feat features.Feature, ddlContent string, queryNames []string, pathBlock string) string {
	var b strings.Builder

	fmt.Fprintf(&b, "Feature:\n")
	fmt.Fprintf(&b, "  op: %s\n", feat.Op)
	fmt.Fprintf(&b, "  path: %s\n", feat.Path)
	fmt.Fprintf(&b, "  table: %s\n", feat.Table)
	fmt.Fprintf(&b, "  public: %v\n", feat.Public)
	fmt.Fprintf(&b, "  desc: %s\n", feat.Desc)

	if ddlContent != "" {
		fmt.Fprintf(&b, "\nDDL for table %s:\n%s\n", feat.Table, ddlContent)
	}

	if len(queryNames) > 0 {
		b.WriteString("\nAvailable sqlc queries:\n")
		for _, name := range queryNames {
			fmt.Fprintf(&b, "  %s\n", name)
		}
	}

	if pathBlock != "" {
		fmt.Fprintf(&b, "\nOpenAPI path block:\n%s\n", pathBlock)
	}

	b.WriteString("\nGenerate a single SSaC service file for this feature.")
	b.WriteString("\nOutput ONLY the SSaC file content. No explanations. No markdown fences.")

	return b.String()
}
