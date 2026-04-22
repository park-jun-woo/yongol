//ff:func feature=manifest type=parser control=iteration dimension=1
//ff:what walkSchemas — components.schemas 매핑의 각 스키마 줄 번호와 property 줄 번호를 색인
package openapi

import "gopkg.in/yaml.v3"

// walkSchemas indexes components.schemas entries and their properties.
func walkSchemas(schemas *yaml.Node, idx *LineIndex) {
	if schemas.Kind != yaml.MappingNode {
		return
	}
	for i := 0; i+1 < len(schemas.Content); i += 2 {
		indexSchemaEntry(schemas.Content[i], schemas.Content[i+1], idx)
	}
}
