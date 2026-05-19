//ff:func feature=features type=util control=iteration dimension=1
//ff:what extractFeatureLines — yaml.Node 트리에서 features 배열 요소별 라인 번호 추출
package features

import "gopkg.in/yaml.v3"

// extractFeatureLines parses the raw YAML into a yaml.Node tree and returns
// the 1-based line number for each element in the features sequence.
func extractFeatureLines(data []byte) []int {
	var root yaml.Node
	if err := yaml.Unmarshal(data, &root); err != nil {
		return nil
	}
	if root.Kind != yaml.DocumentNode || len(root.Content) == 0 {
		return nil
	}
	mapping := root.Content[0]
	if mapping.Kind != yaml.MappingNode {
		return nil
	}
	for i := 0; i+1 < len(mapping.Content); i += 2 {
		key := mapping.Content[i]
		val := mapping.Content[i+1]
		if key.Value == "features" && val.Kind == yaml.SequenceNode {
			return collectSequenceLines(val)
		}
	}
	return nil
}
