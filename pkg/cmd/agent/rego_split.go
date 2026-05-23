//ff:func feature=agent type=helper control=iteration dimension=3
//ff:what extractRegoBlock — Rego allow 블록 추출 (input.action == "operationId" 기준)

package agent

import (
	"fmt"
	"strings"
)

// extractRegoBlock extracts the allow block containing input.action == "operationId".
func extractRegoBlock(content, operationId string) (block string, startLine, endLine int, err error) {
	lines := strings.Split(content, "\n")

	actionPattern := `input.action == "` + operationId + `"`
	actionLine := -1
	for i, l := range lines {
		if strings.Contains(l, actionPattern) {
			actionLine = i
			break
		}
	}
	if actionLine < 0 {
		return "", 0, 0, fmt.Errorf("operationId %q not found in Rego content", operationId)
	}

	blockStart := -1
	for i := actionLine; i >= 0; i-- {
		trimmed := strings.TrimSpace(lines[i])
		if strings.HasPrefix(trimmed, "allow") && strings.Contains(trimmed, "if") && strings.Contains(trimmed, "{") {
			blockStart = i
			break
		}
	}
	if blockStart < 0 {
		return "", 0, 0, fmt.Errorf("could not find 'allow if {' for operationId %q", operationId)
	}

	braceCount := 0
	blockEnd := -1
	for i := blockStart; i < len(lines); i++ {
		for _, c := range lines[i] {
			if c == '{' {
				braceCount++
			} else if c == '}' {
				braceCount--
				if braceCount == 0 {
					blockEnd = i + 1
					break
				}
			}
		}
		if blockEnd >= 0 {
			break
		}
	}
	if blockEnd < 0 {
		return "", 0, 0, fmt.Errorf("could not find closing brace for allow block (operationId %q)", operationId)
	}

	startLine = blockStart
	endLine = blockEnd

	var b strings.Builder
	for i := startLine; i < endLine; i++ {
		b.WriteString(lines[i])
		if i < endLine-1 {
			b.WriteByte('\n')
		}
	}

	return b.String(), startLine, endLine, nil
}
