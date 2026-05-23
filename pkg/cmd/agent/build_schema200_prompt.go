//ff:func feature=agent type=helper control=sequence
//ff:what buildSchema200Prompt — OpenAPI 200 response schema 생성용 프롬프트 구성

package agent

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func buildSchema200Prompt(feat features.Feature, ddlContent string) string {
	var b strings.Builder
	b.WriteString("OpenAPI 200 response schema for this endpoint:\n")
	fmt.Fprintf(&b, "  op: %s\n", feat.Op)
	fmt.Fprintf(&b, "  desc: %s\n", feat.Desc)
	if ddlContent != "" {
		fmt.Fprintf(&b, "\nDDL:\n%s\n", ddlContent)
	}
	b.WriteString("\nAll integers: format: int64. Timestamps: format: date-time.\n")
	b.WriteString("\nOutput ONLY the schema (type, required, properties) in YAML. No wrapping 'schema:' key.")
	return b.String()
}
