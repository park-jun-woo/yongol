//ff:func feature=agent type=helper control=sequence
//ff:what openapi_split — OpenAPI operationId 단위 블록 추출·머지

package agent

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

// extractOpenAPIBlock extracts the path+method block for an operationId
// from openapi.yaml content.
// Returns: extracted yaml string, start line (0-indexed), end line (0-indexed, exclusive).
func extractOpenAPIBlock(content, operationId string) (block string, startLine, endLine int, err error) {
	lines := strings.Split(content, "\n")

	// Find the line containing "operationId: <operationId>"
	opLine := -1
	for i, l := range lines {
		trimmed := strings.TrimSpace(l)
		if trimmed == "operationId: "+operationId {
			opLine = i
			break
		}
	}
	if opLine < 0 {
		return "", 0, 0, fmt.Errorf("operationId %q not found in OpenAPI content", operationId)
	}

	// Walk backwards to find the method line (get:/post:/put:/delete:/patch:)
	// and then the path line above it.
	methodLine := -1
	for i := opLine - 1; i >= 0; i-- {
		trimmed := strings.TrimSpace(lines[i])
		if isHTTPMethod(trimmed) {
			methodLine = i
			break
		}
	}
	if methodLine < 0 {
		return "", 0, 0, fmt.Errorf("could not find method line for operationId %q", operationId)
	}

	// The path line is the parent of the method line — find by indentation.
	methodIndent := countLeadingSpaces(lines[methodLine])
	pathLine := -1
	for i := methodLine - 1; i >= 0; i-- {
		indent := countLeadingSpaces(lines[i])
		trimmed := strings.TrimSpace(lines[i])
		if trimmed == "" {
			continue
		}
		if indent < methodIndent {
			pathLine = i
			break
		}
	}
	if pathLine < 0 {
		return "", 0, 0, fmt.Errorf("could not find path line for operationId %q", operationId)
	}

	// Find the end of this method block.
	// The method block ends when we encounter a line at the same or lesser indent
	// as the method line (i.e., another method or another path), or EOF.
	endIdx := len(lines)
	for i := methodLine + 1; i < len(lines); i++ {
		trimmed := strings.TrimSpace(lines[i])
		if trimmed == "" {
			continue
		}
		indent := countLeadingSpaces(lines[i])
		if indent <= methodIndent {
			endIdx = i
			break
		}
	}

	// Include the path line only as context (first line of block).
	// The block spans from pathLine to endIdx.
	startLine = pathLine
	endLine = endIdx

	var b strings.Builder
	for i := startLine; i < endLine; i++ {
		b.WriteString(lines[i])
		if i < endLine-1 {
			b.WriteByte('\n')
		}
	}

	return b.String(), startLine, endLine, nil
}

// mergeOpenAPIBlock replaces the operationId's path block in the original content.
// Validates yaml structure before merge. Returns modified content or error.
func mergeOpenAPIBlock(originalContent string, startLine, endLine int, fixedBlock string) (string, error) {
	// Validate: the fixedBlock must parse as valid YAML
	var node yaml.Node
	if err := yaml.Unmarshal([]byte(fixedBlock), &node); err != nil {
		return "", fmt.Errorf("fixed OpenAPI block is not valid YAML: %w", err)
	}

	// Validate: must still contain an operationId field
	if !strings.Contains(fixedBlock, "operationId:") {
		return "", fmt.Errorf("fixed OpenAPI block is missing operationId field")
	}

	return spliceLines(originalContent, startLine, endLine, fixedBlock), nil
}

// insertOpenAPIBlock inserts a new path block into openapi.yaml.
// Finds the end of the paths: section (before components: or EOF) and inserts there.
// Validates yaml before insertion.
func insertOpenAPIBlock(originalContent, newBlock string) (string, error) {
	// Validate: the newBlock must parse as valid YAML
	var node yaml.Node
	if err := yaml.Unmarshal([]byte(newBlock), &node); err != nil {
		return "", fmt.Errorf("new OpenAPI block is not valid YAML: %w", err)
	}

	// Validate: must contain an operationId field
	if !strings.Contains(newBlock, "operationId:") {
		return "", fmt.Errorf("new OpenAPI block is missing operationId field")
	}

	lines := strings.Split(originalContent, "\n")

	// Find "paths:" top-level key
	pathsLine := -1
	for i, l := range lines {
		trimmed := strings.TrimSpace(l)
		if trimmed == "paths:" && countLeadingSpaces(l) == 0 {
			pathsLine = i
			break
		}
	}
	if pathsLine < 0 {
		return "", fmt.Errorf("'paths:' section not found in OpenAPI content")
	}

	// Find the end of the paths section: next top-level key (indent 0) or EOF
	insertAt := len(lines)
	for i := pathsLine + 1; i < len(lines); i++ {
		trimmed := strings.TrimSpace(lines[i])
		if trimmed == "" {
			continue
		}
		if countLeadingSpaces(lines[i]) == 0 {
			insertAt = i
			break
		}
	}

	// Insert the new block before the next section (or at EOF)
	newBlock = strings.TrimRight(newBlock, "\n")
	var b strings.Builder
	for i := 0; i < insertAt; i++ {
		b.WriteString(lines[i])
		b.WriteByte('\n')
	}
	b.WriteString(newBlock)
	b.WriteByte('\n')
	for i := insertAt; i < len(lines); i++ {
		b.WriteString(lines[i])
		if i < len(lines)-1 {
			b.WriteByte('\n')
		}
	}

	return b.String(), nil
}

// isHTTPMethod checks if a trimmed line starts with a YAML key for an HTTP method.
func isHTTPMethod(trimmed string) bool {
	methods := []string{"get:", "post:", "put:", "delete:", "patch:", "head:", "options:"}
	lower := strings.ToLower(trimmed)
	for _, m := range methods {
		if lower == m || strings.HasPrefix(lower, m+" ") {
			return true
		}
	}
	return false
}

// countLeadingSpaces returns the number of leading space characters.
func countLeadingSpaces(s string) int {
	n := 0
	for _, c := range s {
		if c == ' ' {
			n++
		} else {
			break
		}
	}
	return n
}
