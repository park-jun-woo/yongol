//ff:func feature=manifest type=parser control=sequence
//ff:what extractBodyConstraints — requestBody에서 필드별 제약조건 추출
package openapi

import "github.com/getkin/kin-openapi/openapi3"

func extractBodyConstraints(body *openapi3.RequestBodyRef, opID string) map[string]FieldConstraint {
	if body.Value == nil || body.Value.Content == nil {
		return nil
	}
	ct := body.Value.Content.Get("application/json")
	if ct == nil || ct.Schema == nil || ct.Schema.Value == nil {
		return nil
	}
	return extractSchemaConstraints(ct.Schema.Value)
}
