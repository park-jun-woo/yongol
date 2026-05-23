//ff:func feature=agent type=helper control=iteration dimension=2
//ff:what extractOpenAPIBlock — OpenAPI operationId 단위 path+method 블록 추출

package agent

import (
	"fmt"
	"strings"
)

// extractOpenAPIBlock extracts the path+method block for an operationId
// from openapi.yaml content.
func extractOpenAPIBlock(content, operationId string) (block string, startLine, endLine int, err error) {
	lines := strings.Split(content, "\n")

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
