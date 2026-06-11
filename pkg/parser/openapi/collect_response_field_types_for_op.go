//ff:func feature=openapi-parse type=parser control=iteration dimension=1
//ff:what collectResponseFieldTypesForOp — 단일 operation 의 2xx 응답 스키마에서 필드 경로별 타입 수집

package openapi

import "github.com/getkin/kin-openapi/openapi3"

// collectResponseFieldTypesForOp fills result[op.OperationID] from the op's
// 2xx JSON response schema.
func collectResponseFieldTypesForOp(result map[string]map[string]FieldTypeInfo, op *openapi3.Operation) {
	if op.OperationID == "" || op.Responses == nil {
		return
	}
	for code, resp := range op.Responses.Map() {
		if len(code) == 0 || code[0] != '2' || resp.Value == nil || resp.Value.Content == nil {
			continue
		}
		ct := resp.Value.Content.Get("application/json")
		if ct == nil || ct.Schema == nil || ct.Schema.Value == nil {
			continue
		}
		fields := collectFieldTypes(ct.Schema.Value)
		if len(fields) > 0 {
			result[op.OperationID] = fields
		}
	}
}
