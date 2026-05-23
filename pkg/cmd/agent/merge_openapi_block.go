//ff:func feature=agent type=helper control=sequence
//ff:what mergeOpenAPIBlock — OpenAPI 블록을 원본 콘텐츠에 머지

package agent

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

// mergeOpenAPIBlock replaces the operationId's path block in the original content.
func mergeOpenAPIBlock(originalContent string, startLine, endLine int, fixedBlock string) (string, error) {
	var node yaml.Node
	if err := yaml.Unmarshal([]byte(fixedBlock), &node); err != nil {
		return "", fmt.Errorf("fixed OpenAPI block is not valid YAML: %w", err)
	}
	if !strings.Contains(fixedBlock, "operationId:") {
		return "", fmt.Errorf("fixed OpenAPI block is missing operationId field")
	}
	return spliceLines(originalContent, startLine, endLine, fixedBlock), nil
}
