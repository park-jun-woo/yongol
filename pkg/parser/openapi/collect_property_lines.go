//ff:func feature=manifest type=parser control=iteration dimension=1
//ff:what collectPropertyLines — properties 매핑의 각 key 노드 줄 번호를 map 으로 수집
package openapi

import "gopkg.in/yaml.v3"

// collectPropertyLines maps each property name in a properties mapping to the
// line of its key node.
func collectPropertyLines(props *yaml.Node) map[string]int {
	out := map[string]int{}
	if props == nil || props.Kind != yaml.MappingNode {
		return out
	}
	for i := 0; i+1 < len(props.Content); i += 2 {
		key := props.Content[i]
		out[key.Value] = key.Line
	}
	return out
}
