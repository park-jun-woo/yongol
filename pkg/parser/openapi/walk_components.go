//ff:func feature=manifest type=parser control=sequence
//ff:what walkComponents — components.schemas 를 찾아 walkSchemas 로 넘김
package openapi

import (
	"gopkg.in/yaml.v3"
)

// walkComponents drills into components.schemas and hands the schemas mapping
// to walkSchemas for indexing.
func walkComponents(comps *yaml.Node, idx *LineIndex) {
	schemasNode := mapValue(comps, "schemas")
	if schemasNode == nil {
		return
	}
	walkSchemas(schemasNode, idx)
}
