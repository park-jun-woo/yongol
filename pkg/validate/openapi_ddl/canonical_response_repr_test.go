//ff:func feature=validate type=test control=iteration dimension=1 topic=openapi-ddl
//ff:what canonicalResponseRepr — XDO-11 발화/침묵, XDO-12 발화, 전략 A/B 경계, 엔티티/비엔티티 구분 서브테스트 디스패치

package openapi_ddl

import "testing"

func TestCanonicalResponseRepr(t *testing.T) {
	for _, st := range []struct {
		name string
		fn   func(*testing.T)
	}{
		{"XDO-11 fires: flat-inline GET vs $ref Update (BUG-131)", subtestTestCanonicalResponseReprXdo11FiresFlatInlineVsRef},
		{"XDO-11 silent: all responses share $ref", subtestTestCanonicalResponseReprXdo11SilentAllShareRef},
		{"XDO-12 fires: consistent but inline (non-SSaC column match)", subtestTestCanonicalResponseReprXdo12FiresConsistentInline},
		{"strategy A: shorthand @response var, divergent components → XDO-11", subtestTestCanonicalResponseReprStrategyADivergentComponents},
		{"non-entity SSaC fields convergence is skipped", subtestTestCanonicalResponseReprNonEntitySkipped},
		{"paginated list response is excluded from grouping", subtestTestCanonicalResponseReprPaginatedListExcluded},
		{"nil OpenAPI doc returns nil", subtestTestCanonicalResponseReprNilDocReturnsNil},
	} {
		t.Run(st.name, st.fn)
	}
}
