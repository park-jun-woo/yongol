//ff:func feature=manifest type=parser control=iteration dimension=1
//ff:what ExtractResponseArrayItemFields — 응답 스키마에서 배열 필드의 항목 프로퍼티 이름을 추출한다
package openapi

import "github.com/getkin/kin-openapi/openapi3"

// ExtractResponseArrayItemFields returns a map of operationId → array field name
// → set of item property names. This is used to determine whether list items
// have an "id" field for React key assignment.
func ExtractResponseArrayItemFields(doc *openapi3.T) map[string]map[string]map[string]bool {
	result := make(map[string]map[string]map[string]bool)
	if doc == nil || doc.Paths == nil {
		return result
	}
	for _, item := range doc.Paths.Map() {
		for _, op := range item.Operations() {
			collectResponseArrayItemFieldsForOp(result, op)
		}
	}
	return result
}

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
		schema := ct.Schema.Value
		fields := make(map[string]map[string]bool)
		for propName, propRef := range schema.Properties {
			if propRef.Value == nil || propRef.Value.Type == nil {
				continue
			}
			if !propRef.Value.Type.Is("array") {
				continue
			}
			items := propRef.Value.Items
			if items == nil || items.Value == nil {
				continue
			}
			itemFields := make(map[string]bool)
			for itemPropName := range items.Value.Properties {
				itemFields[itemPropName] = true
			}
			if len(itemFields) > 0 {
				fields[propName] = itemFields
			}
		}
		if len(fields) > 0 {
			result[op.OperationID] = fields
		}
	}
}
