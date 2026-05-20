//ff:func feature=agent type=helper control=sequence
//ff:what rego_split — Rego allow 블록 단위 추출·머지

package agent

import (
	"fmt"
	"strings"
)

// extractRegoBlock extracts the allow block containing input.action == "operationId".
// Returns: block string, start line (0-indexed), end line (0-indexed, exclusive).
func extractRegoBlock(content, operationId string) (block string, startLine, endLine int, err error) {
	lines := strings.Split(content, "\n")

	// Find the line containing input.action == "operationId"
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

	// Walk backwards to find "allow if {" (the start of this allow block)
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

	// Walk forward from blockStart to find the closing "}"
	braceCount := 0
	blockEnd := -1
	for i := blockStart; i < len(lines); i++ {
		for _, c := range lines[i] {
			if c == '{' {
				braceCount++
			} else if c == '}' {
				braceCount--
				if braceCount == 0 {
					blockEnd = i + 1 // exclusive
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

// mergeRegoBlock replaces the allow block in original content.
// Validates that the fixed block contains "allow if {".
func mergeRegoBlock(originalContent string, startLine, endLine int, fixedBlock string) (string, error) {
	if !strings.Contains(fixedBlock, "allow") || !strings.Contains(fixedBlock, "if") || !strings.Contains(fixedBlock, "{") {
		return "", fmt.Errorf("fixed Rego block is missing 'allow if {' pattern")
	}

	return spliceLines(originalContent, startLine, endLine, fixedBlock), nil
}

// insertRegoBlock appends a new allow block to the end of the rego file.
// Validates that the new block contains "allow if {" pattern.
func insertRegoBlock(originalContent, newBlock string) (string, error) {
	if !strings.Contains(newBlock, "allow") || !strings.Contains(newBlock, "if") || !strings.Contains(newBlock, "{") {
		return "", fmt.Errorf("new Rego block is missing 'allow if {' pattern")
	}

	content := strings.TrimRight(originalContent, "\n")
	newBlock = strings.TrimRight(newBlock, "\n")

	return content + "\n\n" + newBlock + "\n", nil
}
