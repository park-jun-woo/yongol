//ff:func feature=openapi-parse type=parser control=iteration dimension=1
//ff:what 단일 오퍼레이션에서 path 파라미터의 타입을 수집한다

package openapi

import "github.com/getkin/kin-openapi/openapi3"

func collectPathParamTypesForOp(result map[string]map[string]string, op *openapi3.Operation) {
	if op == nil || op.OperationID == "" {
		return
	}
	for _, p := range op.Parameters {
		if p.Value == nil || p.Value.In != "path" {
			continue
		}
		if p.Value.Schema == nil || p.Value.Schema.Value == nil {
			continue
		}
		types := p.Value.Schema.Value.Type.Slice()
		if len(types) == 0 {
			continue
		}
		typStr := types[0]
		if result[op.OperationID] == nil {
			result[op.OperationID] = make(map[string]string)
		}
		result[op.OperationID][p.Value.Name] = typStr
	}
}
