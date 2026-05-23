//ff:func feature=agent type=helper control=iteration dimension=3
//ff:what existingPathBlocks — 기존 openapi.yaml의 path 블록 및 pathToOps 매핑 읽기

package agent

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func existingPathBlocks(specsDir string) (map[string]any, map[string][]string) {
	data, err := os.ReadFile(filepath.Join(specsDir, "api", "openapi.yaml"))
	if err != nil {
		return nil, nil
	}
	var doc struct {
		Paths map[string]any `yaml:"paths"`
	}
	if err := yaml.Unmarshal(data, &doc); err != nil || doc.Paths == nil {
		return nil, nil
	}
	pathToOps := make(map[string][]string)
	for pathKey, methods := range doc.Paths {
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
				pathToOps[pathKey] = appendUnique(pathToOps[pathKey], fmt.Sprintf("%v", opID))
			}
		}
	}
	return doc.Paths, pathToOps
}
