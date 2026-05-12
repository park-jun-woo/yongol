//ff:func feature=openapi-parse type=parser control=iteration dimension=1
//ff:what collectResponseArrayItemFieldsForOp — 단일 operation 의 응답 스키마에서 배열 항목 프로퍼티 수집

package openapi

import "github.com/getkin/kin-openapi/openapi3"

func collectResponseArrayItemFieldsForOp(result map[string]map[string]map[string]bool, op *openapi3.Operation) {
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
		fields := extractArrayItemFields(ct.Schema.Value)
		if len(fields) > 0 {
			result[op.OperationID] = fields
		}
	}
}
