//ff:func feature=manifest type=parser control=iteration dimension=1
//ff:what indexPathItem — 단일 PathItem 의 operation 들을 색인
package openapi

import "gopkg.in/yaml.v3"

// indexPathItem records the path key's line and walks its method operations.
// pathKey / pathItem must come as a consecutive key/value pair from a
// paths MappingNode.
func indexPathItem(pathKey, pathItem *yaml.Node, idx *LineIndex) {
	idx.Paths[pathKey.Value] = pathKey.Line
	if pathItem.Kind != yaml.MappingNode {
		return
	}
	for j := 0; j+1 < len(pathItem.Content); j += 2 {
		methodKey := pathItem.Content[j]
		opNode := pathItem.Content[j+1]
		if !httpMethods[methodKey.Value] || opNode.Kind != yaml.MappingNode {
			continue
		}
		indexOperation(opNode, idx)
	}
}
