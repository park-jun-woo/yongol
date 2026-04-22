//ff:func feature=manifest type=parser control=sequence
//ff:what indexRequestBody — requestBody 아래 schema.properties 의 각 필드 줄 번호를 색인
package openapi

import "gopkg.in/yaml.v3"

// indexRequestBody records RequestFields lines for the given operationId by
// reading schema.properties from the requestBody node.
func indexRequestBody(rb *yaml.Node, opID string, idx *LineIndex) {
	props := schemaPropsOfBody(rb)
	if props == nil {
		return
	}
	idx.RequestFields[opID] = collectPropertyLines(props)
}
