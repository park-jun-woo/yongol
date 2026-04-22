//ff:func feature=manifest type=parser control=sequence
//ff:what schemaPropsOfBody — body → content → application/json → schema → properties 노드 반환
package openapi

import "gopkg.in/yaml.v3"

// schemaPropsOfBody walks down body → content → application/json → schema →
// properties and returns the properties mapping node, or nil.
func schemaPropsOfBody(body *yaml.Node) *yaml.Node {
	content := mapValue(body, "content")
	if content == nil {
		return nil
	}
	mediaType := mapValue(content, "application/json")
	if mediaType == nil {
		return nil
	}
	schema := mapValue(mediaType, "schema")
	if schema == nil {
		return nil
	}
	props := mapValue(schema, "properties")
	if props == nil || props.Kind != yaml.MappingNode {
		return nil
	}
	return props
}
