//ff:func feature=agent type=helper control=iteration dimension=1
//ff:what buildUserPrompt — feature desc + 파일 내용 + 진단으로 user prompt 구성

package agent

import (
	"fmt"
	"strings"
)

// buildUserPrompt assembles the user prompt from feature desc, file content, and diagnostics.
func buildUserPrompt(desc, path, filename, content string, messages []string) string {
	var b strings.Builder
	if desc != "" {
		fmt.Fprintf(&b, "Feature: %s\nPath: %s\n\n", desc, path)
	}
	fmt.Fprintf(&b, "Current file (%s):\n%s\n\nValidate errors:\n", filename, content)
	for _, m := range messages {
		b.WriteString(m)
		b.WriteByte('\n')
	}
	b.WriteString("\nFix the file. Output ONLY the corrected file content.")
	return b.String()
}
