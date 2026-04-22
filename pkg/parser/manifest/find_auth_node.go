//ff:func feature=projectconfig type=parser control=sequence
//ff:what FindAuthNode — manifest.yaml 원본에서 backend.auth yaml.Node 를 찾아 반환
package manifest

import (
	"gopkg.in/yaml.v3"
)

// FindAuthNode unmarshals the raw manifest.yaml bytes and returns the
// backend.auth yaml.Node, or nil if absent/invalid. Exposed so validator
// packages (e.g. pkg/validate/manifest) can inspect keys that are not
// present in the parsed Auth struct — the `secret` literal rejected by
// SEC-401 is the primary use case.
func FindAuthNode(data []byte) *yaml.Node {
	var root yaml.Node
	if err := yaml.Unmarshal(data, &root); err != nil {
		return nil
	}
	doc := &root
	if doc.Kind == yaml.DocumentNode && len(doc.Content) > 0 {
		doc = doc.Content[0]
	}
	if doc.Kind != yaml.MappingNode {
		return nil
	}
	backend := mappingValue(doc, "backend")
	if backend == nil {
		return nil
	}
	return mappingValue(backend, "auth")
}
