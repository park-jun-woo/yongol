//ff:func feature=manifest type=parser control=iteration dimension=1
//ff:what mapValue — yaml.MappingNode 에서 주어진 key 의 value 노드를 반환
package openapi

import "gopkg.in/yaml.v3"

// mapValue returns the value node for key k inside a mapping node, or nil.
func mapValue(m *yaml.Node, k string) *yaml.Node {
	if m == nil || m.Kind != yaml.MappingNode {
		return nil
	}
	for i := 0; i+1 < len(m.Content); i += 2 {
		if m.Content[i].Value == k {
			return m.Content[i+1]
		}
	}
	return nil
}
