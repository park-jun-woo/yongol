//ff:func feature=agent type=helper control=sequence
//ff:what buildRetryPrompt — kin-openapi 검증 실패 후 재시도 프롬프트 구성

package agent

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func buildRetryPrompt(feat features.Feature, ddlContent, prevError string, relativeLine int) string {
	var b strings.Builder

	fmt.Fprintf(&b, "Feature:\n")
	fmt.Fprintf(&b, "  op: %s\n", feat.Op)
	fmt.Fprintf(&b, "  path: %s\n", feat.Path)
	fmt.Fprintf(&b, "  table: %s\n", feat.Table)
	fmt.Fprintf(&b, "  public: %v\n", feat.Public)
	fmt.Fprintf(&b, "  desc: %s\n", feat.Desc)

	if ddlContent != "" {
		fmt.Fprintf(&b, "\nTable DDL:\n%s\n", ddlContent)
	}

	fmt.Fprintf(&b, "\nPrevious attempt had this error:\n%s\n", prevError)
	if relativeLine >= 0 {
		fmt.Fprintf(&b, "\nThe error is near line %d of your path block.\n", relativeLine)
	}
	b.WriteString("\nWrite the corrected OpenAPI path block.")
	b.WriteString("\nThe output must be a valid YAML fragment starting with the path key (e.g. /resources/{id}:).")
	b.WriteString("\nInclude operationId, request/response schemas.")
	b.WriteString("\nUse 2-space indentation. No surrounding 'paths:' key. No markdown fences.")

	return b.String()
}
