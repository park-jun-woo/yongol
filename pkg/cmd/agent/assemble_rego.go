//ff:func feature=agent type=helper control=iteration dimension=1
//ff:what assembleRego — Rego 룰 블록들을 완전한 정책 파일로 조합

package agent

import "strings"

// assembleRego combines rule blocks into a complete Rego policy file.
func assembleRego(ruleBlocks []string) string {
	var b strings.Builder

	b.WriteString("package authz\n\nimport rego.v1\n\ndefault allow := false\n")

	for _, block := range ruleBlocks {
		cleaned := cleanRegoBlock(block)
		if cleaned != "" {
			b.WriteString("\n")
			b.WriteString(cleaned)
			b.WriteString("\n")
		}
	}

	return b.String()
}
