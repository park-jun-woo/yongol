//ff:func feature=rule type=loader control=iteration dimension=1
//ff:what populateOpParams — 단일 operation의 param, request를 Ground에 등록
package ground

import (
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

func populateOpParams(g *rule.Ground, op *openapi3.Operation) {
	opID := op.OperationID
	params := make(rule.StringSet)
	for _, p := range op.Parameters {
		if p.Value != nil {
			params[p.Value.Name] = true
			// Register param Go type for XFS-73 (request.* param type resolution).
			if p.Value.Schema != nil && p.Value.Schema.Value != nil {
				sv := p.Value.Schema.Value
				if sv.Type != nil && len(*sv.Type) > 0 {
					goType := resolveOAPIParamGoType((*sv.Type)[0], sv.Format)
					if goType != "" {
						g.Types["OpenAPI.paramType."+opID+"."+p.Value.Name] = goType
					}
				}
			}
		}
	}
	g.Lookup["OpenAPI.param."+opID] = params

	if op.RequestBody != nil && op.RequestBody.Value != nil {
		reqFields := extractRequestFields(op.RequestBody.Value)
		if len(reqFields) > 0 {
			g.Lookup["OpenAPI.request."+opID] = reqFields
		}
	}
}
