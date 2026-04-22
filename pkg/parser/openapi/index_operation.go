//ff:func feature=manifest type=parser control=sequence
//ff:what indexOperation — 단일 operation 의 operationId, request/response 필드를 색인
package openapi

import "gopkg.in/yaml.v3"

// indexOperation extracts operationId and properties for request/response.
func indexOperation(op *yaml.Node, idx *LineIndex) {
	opIDNode := mapValue(op, "operationId")
	if opIDNode == nil || opIDNode.Value == "" {
		return
	}
	opID := opIDNode.Value
	// operationId 키가 등장한 줄을 operation 의 대표 line 으로 사용.
	if k := mapKey(op, "operationId"); k != nil {
		idx.Operations[opID] = k.Line
	} else {
		idx.Operations[opID] = opIDNode.Line
	}

	if rb := mapValue(op, "requestBody"); rb != nil {
		indexRequestBody(rb, opID, idx)
	}
	if resps := mapValue(op, "responses"); resps != nil && resps.Kind == yaml.MappingNode {
		indexFirst2xxResponse(resps, opID, idx)
	}
}
