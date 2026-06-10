//ff:func feature=gen-react type=util control=iteration dimension=2
//ff:what findRefreshOp — 선언된 frontend.auth.refresh_op 의 operationId 로 refreshPlan 구성

package react

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

// findRefreshOp locates the operation declared as frontend.auth.refresh_op
// and builds its refreshPlan. found is false when no operation carries that
// operationId (validate's XON-60 already diagnoses this as an ERROR; the
// emitter degrades to the no-refresh client instead of failing generate).
func findRefreshOp(doc *openapi3.T, fa *manifest.FrontendAuth) (*refreshPlan, bool) {
	for path, pi := range doc.Paths.Map() {
		for method, op := range pi.Operations() {
			if op == nil || op.OperationID != fa.RefreshOp {
				continue
			}
			return &refreshPlan{
				opID:         op.OperationID,
				method:       method,
				path:         path,
				tokenField:   fa.TokenField,
				refreshField: fa.RefreshField,
				bodyKey:      refreshBodyKey(op, fa.RefreshField),
			}, true
		}
	}
	return nil, false
}
