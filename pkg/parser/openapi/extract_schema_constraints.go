//ff:func feature=manifest type=parser control=iteration dimension=1
//ff:what extractSchemaConstraints — 스키마 properties에서 필드별 제약조건 추출
package openapi

import "github.com/getkin/kin-openapi/openapi3"

func extractSchemaConstraints(schema *openapi3.Schema) map[string]FieldConstraint {
	fields := make(map[string]FieldConstraint)
	reqSet := make(map[string]bool, len(schema.Required))
	for _, r := range schema.Required {
		reqSet[r] = true
	}
	for name, ref := range schema.Properties {
		if ref.Value == nil {
			continue
		}
		fields[name] = buildFieldConstraint(ref.Value, reqSet[name])
	}
	return fields
}
