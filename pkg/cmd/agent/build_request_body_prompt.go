//ff:func feature=agent type=helper control=sequence
//ff:what buildRequestBodyPrompt — OpenAPI requestBody 생성용 프롬프트 구성

package agent

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func buildRequestBodyPrompt(feat features.Feature, ddlContent string) string {
	var b strings.Builder
	b.WriteString("OpenAPI requestBody for this endpoint:\n")
	fmt.Fprintf(&b, "  op: %s\n", feat.Op)
	fmt.Fprintf(&b, "  path: %s\n", feat.Path)
	fmt.Fprintf(&b, "  desc: %s\n", feat.Desc)
	if ddlContent != "" {
		fmt.Fprintf(&b, "\nDDL:\n%s\n", ddlContent)
	}
	b.WriteString("\nOutput \"none\" if no requestBody needed, or the requestBody YAML block.")
	return b.String()
}
