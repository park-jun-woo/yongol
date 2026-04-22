//ff:func feature=manifest type=parser control=sequence
//ff:what BuildLineIndex — openapi.yaml 을 yaml.v3 로 재파싱해 LineIndex 를 구축
package openapi

import (
	"os"

	"gopkg.in/yaml.v3"
)

// BuildLineIndex parses the given openapi.yaml with yaml.v3 in node mode and
// returns a LineIndex. Errors are non-fatal: a partially populated index is
// still returned (any field may be empty). 호출자는 line 0 을 "미상" 으로 간주.
func BuildLineIndex(path string) (*LineIndex, error) {
	idx := &LineIndex{
		File:             path,
		Operations:       map[string]int{},
		RequestFields:    map[string]map[string]int{},
		ResponseFields:   map[string]map[string]int{},
		Schemas:          map[string]int{},
		SchemaProperties: map[string]map[string]int{},
		Paths:            map[string]int{},
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return idx, err
	}
	var root yaml.Node
	if err := yaml.Unmarshal(data, &root); err != nil {
		return idx, err
	}
	if root.Kind != yaml.DocumentNode || len(root.Content) == 0 {
		return idx, nil
	}
	doc := root.Content[0]
	if doc.Kind != yaml.MappingNode {
		return idx, nil
	}

	if pathsNode := mapValue(doc, "paths"); pathsNode != nil {
		walkPaths(pathsNode, idx)
	}
	if compsNode := mapValue(doc, "components"); compsNode != nil {
		walkComponents(compsNode, idx)
	}
	return idx, nil
}
