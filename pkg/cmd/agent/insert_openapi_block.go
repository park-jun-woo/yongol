//ff:func feature=agent type=helper control=iteration dimension=1
//ff:what insertOpenAPIBlock — 새 OpenAPI path 블록을 paths 섹션 끝에 삽입

package agent

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

// insertOpenAPIBlock inserts a new path block into openapi.yaml.
func insertOpenAPIBlock(originalContent, newBlock string) (string, error) {
	var node yaml.Node
	if err := yaml.Unmarshal([]byte(newBlock), &node); err != nil {
		return "", fmt.Errorf("new OpenAPI block is not valid YAML: %w", err)
	}
	if !strings.Contains(newBlock, "operationId:") {
		return "", fmt.Errorf("new OpenAPI block is missing operationId field")
	}

	lines := strings.Split(originalContent, "\n")

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
