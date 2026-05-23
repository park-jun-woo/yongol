//ff:func feature=agent type=helper control=iteration dimension=3
//ff:what existingOperationIDs — 기존 openapi.yaml에서 operationId 수집

package agent

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func existingOperationIDs(specsDir string) map[string]bool {
	data, err := os.ReadFile(filepath.Join(specsDir, "api", "openapi.yaml"))
	if err != nil {
		return nil
	}
	var doc struct {
		Paths map[string]any `yaml:"paths"`
	}
	if err := yaml.Unmarshal(data, &doc); err != nil || doc.Paths == nil {
		return nil
	}
	ids := make(map[string]bool)
	for _, methods := range doc.Paths {
		methodMap, ok := methods.(map[string]any)
		if !ok {
			continue
		}
		for _, detail := range methodMap {
			detailMap, ok := detail.(map[string]any)
			if !ok {
				continue
			}
			if opID, ok := detailMap["operationId"]; ok {
				ids[fmt.Sprintf("%v", opID)] = true
			}
		}
	}
	return ids
}
