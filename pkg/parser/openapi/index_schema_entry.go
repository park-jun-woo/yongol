//ff:func feature=manifest type=parser control=sequence
//ff:what indexSchemaEntry — 단일 components.schemas 항목의 줄 번호와 property 줄 번호를 색인
package openapi

import "gopkg.in/yaml.v3"

// indexSchemaEntry records the schema name's line and its property lines.
// nameKey / schemaNode must come from a consecutive key/value pair of the
// components.schemas MappingNode.
func indexSchemaEntry(nameKey, schemaNode *yaml.Node, idx *LineIndex) {
	idx.Schemas[nameKey.Value] = nameKey.Line
	if schemaNode.Kind != yaml.MappingNode {
		return
	}
	props := mapValue(schemaNode, "properties")
	if props == nil {
		return
	}
	idx.SchemaProperties[nameKey.Value] = collectPropertyLines(props)
}
