//ff:func feature=agent type=helper control=iteration dimension=1
//ff:what cleanRegoBlock — Rego 블록에서 중복 package/import/default 라인 제거

package agent

import "strings"

// cleanRegoBlock strips duplicate package/import/default lines from a LLM-generated rego block.
func cleanRegoBlock(block string) string {
	block = strings.TrimSpace(block)
	lines := strings.Split(block, "\n")
	var filtered []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "package ") {
			continue
		}
		if strings.HasPrefix(trimmed, "import ") {
			continue
		}
		if strings.HasPrefix(trimmed, "default allow") {
			continue
		}
		filtered = append(filtered, line)
	}
	return strings.TrimSpace(strings.Join(filtered, "\n"))
}
