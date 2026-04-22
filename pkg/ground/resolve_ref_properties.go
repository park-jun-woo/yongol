//ff:func feature=rule type=util control=iteration dimension=1
//ff:what resolveRefProperties — OpenAPI response schema의 $ref를 풀어 실제 필드명 수집
package ground

import "github.com/getkin/kin-openapi/openapi3"

func resolveRefProperties(schema *openapi3.Schema) []string {
	var fields []string
	for _, ref := range schema.Properties {
		if ref.Value == nil {
			continue
		}
		fields = append(fields, resolveSchemaRef(ref.Value)...)
	}
	return fields
}
