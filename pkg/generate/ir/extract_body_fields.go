//ff:func feature=gen-ir type=util control=iteration dimension=1
//ff:what extractBodyFields -- OpenAPI request body 스키마에서 BodyFieldMeta 추출

package ir

import "github.com/getkin/kin-openapi/openapi3"

// extractBodyFields extracts request body property metadata from an OpenAPI
// operation.
func extractBodyFields(op *openapi3.Operation) []BodyFieldMeta {
	if op.RequestBody == nil || op.RequestBody.Value == nil {
		return nil
	}
	mt := op.RequestBody.Value.Content.Get("application/json")
	if mt == nil || mt.Schema == nil || mt.Schema.Value == nil {
		return nil
	}
	schema := mt.Schema.Value
	requiredSet := make(map[string]bool, len(schema.Required))
	for _, r := range schema.Required {
		requiredSet[r] = true
	}
	var fields []BodyFieldMeta
	for name, propRef := range schema.Properties {
		if propRef == nil || propRef.Value == nil {
			continue
		}
		fields = append(fields, BodyFieldMeta{
			Name:     name,
			Required: requiredSet[name],
			Format:   propRef.Value.Format,
		})
	}
	return fields
}
