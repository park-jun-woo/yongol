//ff:func feature=projectconfig type=parser control=iteration dimension=1
//ff:what mappingValue — yaml.MappingNode 에서 주어진 key 의 value 노드를 반환
package manifest

import (
	"gopkg.in/yaml.v3"
)

// mappingValue returns the value node for the given key in a mapping node,
// or nil if not found / not a mapping.
func mappingValue(n *yaml.Node, key string) *yaml.Node {
	if n == nil || n.Kind != yaml.MappingNode {
		return nil
	}
	for i := 0; i+1 < len(n.Content); i += 2 {
		if n.Content[i].Value == key {
			return n.Content[i+1]
		}
	}
	return nil
}
