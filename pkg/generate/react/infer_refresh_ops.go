//ff:func feature=gen-react type=util control=iteration dimension=2
//ff:what inferRefreshOps — 2xx 에 token+refresh 필드를 모두 가진 capture 미선언 op 후보 수집

package react

import (
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

// inferRefreshOps collects the structural refresh-op candidates: operations
// whose 2xx JSON response declares both frontend.auth.token_field and
// refresh_field, excluding ops bound by an STML data-capture declaration
// (login-style producers) and ops with path parameters (the 401 hook cannot
// supply them). Sorted by operationId so the ambiguity diagnostic and the
// single-candidate pick are deterministic.
func inferRefreshOps(doc *openapi3.T, fa *manifest.FrontendAuth, captured map[string]bool) []refreshPlan {
	var out []refreshPlan
	for path, pi := range doc.Paths.Map() {
		if strings.Contains(path, "{") {
			continue
		}
		for method, op := range pi.Operations() {
			if op == nil || op.OperationID == "" || captured[op.OperationID] {
				continue
			}
			props := op2xxResponseProps(op)
			if !props[fa.TokenField] || !props[fa.RefreshField] {
				continue
			}
			out = append(out, refreshPlan{
				opID:         op.OperationID,
				method:       method,
				path:         path,
				tokenField:   fa.TokenField,
				refreshField: fa.RefreshField,
				bodyKey:      refreshBodyKey(op, fa.RefreshField),
			})
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].opID < out[j].opID })
	return out
}
