//ff:func feature=validate type=util control=iteration dimension=2 topic=config-check
//ff:what collectFrontendAuthProps — 전체 op 의 2xx 응답 프로퍼티 합집합 + refresh_op 의 프로퍼티 set 수집

package openapi_manifest

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// collectFrontendAuthProps walks every operation of doc and returns the
// union of all 2xx JSON response property names, plus the property set of
// the operation whose operationId equals refreshOp (and whether such an
// operation exists). Used by XON-60.
func collectFrontendAuthProps(doc *openapi3.T, refreshOp string) (allProps, refreshOpProps map[string]bool, refreshOpFound bool) {
	allProps = map[string]bool{}
	refreshOpProps = map[string]bool{}
	if doc == nil || doc.Paths == nil {
		return allProps, refreshOpProps, false
	}
	for _, pi := range doc.Paths.Map() {
		for _, op := range pi.Operations() {
			props := op2xxPropertySet(op)
			for name := range props {
				allProps[name] = true
			}
			if op != nil && refreshOp != "" && op.OperationID == refreshOp {
				refreshOpFound = true
				refreshOpProps = props
			}
		}
	}
	return allProps, refreshOpProps, refreshOpFound
}
