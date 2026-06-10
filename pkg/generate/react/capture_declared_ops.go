//ff:func feature=gen-react type=util control=iteration dimension=2
//ff:what captureDeclaredOps — STML data-capture 가 선언된 action 의 operationId set

package react

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// captureDeclaredOps returns the operationIds bound by an STML data-capture
// declaration. Those ops are login-style token producers and are excluded
// from structural refresh-op inference (a refresh op is by definition the
// token-yielding op no page captures explicitly).
func captureDeclaredOps(pages []stml.PageSpec) map[string]bool {
	ops := map[string]bool{}
	for _, p := range pages {
		for _, a := range p.Actions {
			if len(a.Captures) > 0 && a.OperationID != "" {
				ops[a.OperationID] = true
			}
		}
	}
	return ops
}
