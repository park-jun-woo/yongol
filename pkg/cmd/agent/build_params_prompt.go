//ff:func feature=agent type=helper control=sequence
//ff:what buildParamsPrompt — OpenAPI parameters 생성용 프롬프트 구성

package agent

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func buildParamsPrompt(feat features.Feature) string {
	var b strings.Builder
	b.WriteString("OpenAPI parameters for this endpoint:\n")
	fmt.Fprintf(&b, "  op: %s\n", feat.Op)
	fmt.Fprintf(&b, "  path: %s\n", feat.Path)
	fmt.Fprintf(&b, "  desc: %s\n", feat.Desc)
	b.WriteString("\nRules:\n")
	b.WriteString("- Path parameters need matching parameters[] entry\n")
	b.WriteString("- All integers: type: integer, format: int64\n")
	b.WriteString("\nOutput ONLY the parameters array in YAML. Output \"none\" if no parameters.")
	return b.String()
}
