//ff:func feature=manifest type=parser control=iteration dimension=1
//ff:what extractResponseFields — operation의 2xx 응답에서 필드 제약조건 추출
package openapi

import "github.com/getkin/kin-openapi/openapi3"

func extractResponseFields(op *openapi3.Operation) map[string]FieldConstraint {
	for code, resp := range op.Responses.Map() {
		if len(code) == 0 || code[0] != '2' || resp.Value == nil || resp.Value.Content == nil {
			continue
		}
		ct := resp.Value.Content.Get("application/json")
		if ct == nil || ct.Schema == nil || ct.Schema.Value == nil {
			continue
		}
		return extractSchemaConstraints(ct.Schema.Value)
	}
	return nil
}
