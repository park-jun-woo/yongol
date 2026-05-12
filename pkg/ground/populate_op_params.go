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
		if p.Value == nil {
			continue
		}
		params[p.Value.Name] = true
		registerParamGoType(g, opID, p.Value)
	}
	g.Lookup["OpenAPI.param."+opID] = params

	if op.RequestBody != nil && op.RequestBody.Value != nil {
		reqFields := extractRequestFields(op.RequestBody.Value)
		if len(reqFields) > 0 {
			g.Lookup["OpenAPI.request."+opID] = reqFields
		}
	}
}
