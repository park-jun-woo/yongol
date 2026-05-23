//ff:func feature=agent type=helper control=iteration dimension=1
//ff:what buildBlockUserPrompt — 단일 블록 수정용 user prompt 구성

package agent

import (
	"fmt"
	"strings"
)

// buildBlockUserPrompt assembles the user prompt for a single block fix.
func buildBlockUserPrompt(desc, path, filename, opID, block string, messages []string) string {
	var b strings.Builder
	if desc != "" {
		fmt.Fprintf(&b, "Feature: %s\nPath: %s\n\n", desc, path)
	}
	fmt.Fprintf(&b, "OperationId: %s\nFile: %s\n\nCurrent block:\n%s\n\nValidate errors:\n", opID, filename, block)
	for _, m := range messages {
		b.WriteString(m)
		b.WriteByte('\n')
	}
	b.WriteString("\nFix ONLY this block. Output ONLY the corrected block content. Do not add surrounding file content.")
	return b.String()
}
