//ff:func feature=agent type=helper control=iteration dimension=3
//ff:what extractPathBlockForOp — openapi.yaml에서 operationId의 path 블록 추출

package agent

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

func extractPathBlockForOp(openapiContent, operationId string) string {
	var doc struct {
		Paths map[string]any `yaml:"paths"`
	}
	if err := yaml.Unmarshal([]byte(openapiContent), &doc); err != nil || doc.Paths == nil {
		return ""
	}

	for pathKey, methods := range doc.Paths {
		methodMap, ok := methods.(map[string]any)
		if !ok {
			continue
		}
		for method, detail := range methodMap {
			detailMap, ok := detail.(map[string]any)
			if !ok {
				continue
			}
			if opID, ok := detailMap["operationId"]; ok && fmt.Sprintf("%v", opID) == operationId {
				block := map[string]any{pathKey: map[string]any{method: detail}}
				out, err := yaml.Marshal(block)
				if err != nil {
					return ""
				}
				return string(out)
			}
		}
	}
	return ""
}
