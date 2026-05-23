//ff:func feature=agent type=helper control=iteration dimension=1
//ff:what buildRegoUserPrompt — Rego 정책 생성용 user prompt 구성

package agent

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

// buildRegoUserPrompt builds the user prompt for a batch of features.
func buildRegoUserPrompt(feats []features.Feature) string {
	var b strings.Builder

	b.WriteString("Non-public features requiring authorization:\n\n")
	for _, f := range feats {
		resource := f.Table
		if resource == "" {
			resource = domainFromPath(f.Path)
		}
		fmt.Fprintf(&b, "  - op: %s, resource: %s\n", f.Op, resource)
	}

	b.WriteString("\nGenerate OPA Rego allow rules for these features.")
	b.WriteString("\nUse 'allow if { ... }' syntax with input.action == \"<operationId>\" checks.")
	b.WriteString("\nOutput ONLY the allow rule blocks. No package declaration. No import. No markdown fences.")

	return b.String()
}
