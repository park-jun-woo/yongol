//ff:func feature=validate-contract type=util control=sequence
//ff:what expectedSignature — operationId 기준 기대 FuncSignature 계산 (Ground 기반)

package contract

import (
	"github.com/park-jun-woo/yongol/pkg/contract"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

// expectedSignature builds the FuncSignature that yongol generates for
// an HTTP service handler keyed by operationId. The oapi-codegen + SSaC
// generator emits:
//
//	func (server *Server) <OpID>(ctx context.Context,
//	    request api.<OpID>RequestObject) (api.<OpID>ResponseObject, error)
//
// ok is false when the operationId is not present in the OpenAPI Ground
// lookup, in which case callers should emit a dedicated "removed opID"
// diagnostic rather than attempting a field-by-field comparison.
func expectedSignature(g *rule.Ground, opID string) (contract.FuncSignature, bool) {
	if g == nil || opID == "" {
		return contract.FuncSignature{}, false
	}
	opIDs, ok := g.Lookup["OpenAPI.operationId"]
	if !ok || !opIDs[opID] {
		return contract.FuncSignature{}, false
	}
	sig := contract.FuncSignature{
		Name: opID,
		Params: []contract.FuncParam{
			{Name: "ctx", Type: "context.Context"},
			{Name: "request", Type: "api." + opID + "RequestObject"},
		},
		Returns: []string{"api." + opID + "ResponseObject", "error"},
		HasErr:  true,
	}
	return sig, true
}
