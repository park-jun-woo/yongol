//ff:func feature=rule type=util control=iteration dimension=1
//ff:what collectSchemaFieldNames — schema.Properties 의 top-level 필드명 슬라이스 반환
package ground

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// collectSchemaFieldNames returns the top-level property names of a JSON
// schema as an unordered string slice.
func collectSchemaFieldNames(schema *openapi3.Schema) []string {
	var fields []string
	for name := range schema.Properties {
		fields = append(fields, name)
	}
	return fields
}
