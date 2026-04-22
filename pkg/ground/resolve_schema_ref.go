//ff:func feature=rule type=util control=sequence
//ff:what resolveSchemaRef — 단일 schema property의 $ref를 풀어 내부 필드명 반환
package ground

import "github.com/getkin/kin-openapi/openapi3"

func resolveSchemaRef(prop *openapi3.Schema) []string {
	if len(prop.Properties) > 0 {
		var fields []string
		for inner := range prop.Properties {
			fields = append(fields, inner)
		}
		return fields
	}
	if prop.Type.Is("array") && prop.Items != nil && prop.Items.Value != nil {
		var fields []string
		for inner := range prop.Items.Value.Properties {
			fields = append(fields, inner)
		}
		return fields
	}
	return nil
}
