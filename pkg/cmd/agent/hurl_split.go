//ff:func feature=agent type=helper control=sequence
//ff:what hurl_split — Hurl 요청 블록 단위 추출·머지

package agent

import (
	"fmt"
	"strings"
)

// extractHurlBlock extracts the request block starting with "# OperationId" comment.
// Block ends at next "# " comment or EOF.
// Returns: block string, start line (0-indexed), end line (0-indexed, exclusive).
func extractHurlBlock(content, operationId string) (block string, startLine, endLine int, err error) {
	lines := strings.Split(content, "\n")

	// Find the line "# <operationId>"
	commentPrefix := "# " + operationId
	blockStart := -1
	for i, l := range lines {
		trimmed := strings.TrimSpace(l)
		if trimmed == commentPrefix {
			blockStart = i
			break
		}
	}
	if blockStart < 0 {
		return "", 0, 0, fmt.Errorf("operationId %q not found as Hurl comment (# %s)", operationId, operationId)
	}

	// Find the end: next "# " comment line or EOF
	blockEnd := len(lines)
	for i := blockStart + 1; i < len(lines); i++ {
		trimmed := strings.TrimSpace(lines[i])
		if strings.HasPrefix(trimmed, "# ") && !strings.HasPrefix(trimmed, "# @") {
			// Next operation block starts here
			blockEnd = i
			break
		}
	}

	// Trim trailing empty lines from block
	for blockEnd > blockStart+1 && strings.TrimSpace(lines[blockEnd-1]) == "" {
		blockEnd--
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

// mergeHurlBlock replaces the request block in original content.
// Validates that the fixed block contains an HTTP method line.
func mergeHurlBlock(originalContent string, startLine, endLine int, fixedBlock string) (string, error) {
	if !containsHTTPMethodLine(fixedBlock) {
		return "", fmt.Errorf("fixed Hurl block is missing HTTP method line (GET/POST/PUT/DELETE/PATCH)")
	}

	return spliceLines(originalContent, startLine, endLine, fixedBlock), nil
}

// insertHurlBlock appends a new request block to the end of the hurl file.
// Validates that the new block contains an HTTP method line.
func insertHurlBlock(originalContent, newBlock string) (string, error) {
	if !containsHTTPMethodLine(newBlock) {
		return "", fmt.Errorf("new Hurl block is missing HTTP method line (GET/POST/PUT/DELETE/PATCH)")
	}

	content := strings.TrimRight(originalContent, "\n")
	newBlock = strings.TrimRight(newBlock, "\n")

	return content + "\n\n" + newBlock + "\n", nil
}

// containsHTTPMethodLine checks if the text contains a line starting with an HTTP method.
func containsHTTPMethodLine(s string) bool {
	methods := []string{"GET ", "POST ", "PUT ", "DELETE ", "PATCH ", "HEAD ", "OPTIONS "}
	for _, line := range strings.Split(s, "\n") {
		trimmed := strings.TrimSpace(line)
		for _, m := range methods {
			if strings.HasPrefix(trimmed, m) {
				return true
			}
		}
	}
	return false
}
