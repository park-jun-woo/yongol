//ff:func feature=agent type=helper control=sequence
//ff:what mergeRegoBlock — Rego allow 블록을 원본 콘텐츠에 머지

package agent

import (
	"fmt"
	"strings"
)

// mergeRegoBlock replaces the allow block in original content.
func mergeRegoBlock(originalContent string, startLine, endLine int, fixedBlock string) (string, error) {
	if !strings.Contains(fixedBlock, "allow") || !strings.Contains(fixedBlock, "if") || !strings.Contains(fixedBlock, "{") {
		return "", fmt.Errorf("fixed Rego block is missing 'allow if {' pattern")
	}
	return spliceLines(originalContent, startLine, endLine, fixedBlock), nil
}
