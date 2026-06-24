//ff:func feature=gen-react type=util control=iteration dimension=1
//ff:what resolveLayoutLogoutOps — bearer 모드에서 valueless data-logout의 OperationID를 OpenAPI logout op으로 자동 해석

package react

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/validate/stml_openapi"
)

// resolveLayoutLogoutOps fills in the OperationID for layouts that declare a
// valueless data-logout (OperationID == "") in bearer auth mode. It queries
// the OpenAPI document for a logout-like operation (FindLogoutOp) and wires
// it automatically so the generated logout handler calls the server endpoint.
// LayoutSpec.Logout is a pointer, so mutation propagates to downstream
// consumers (writeLayoutsTSX, renderLayoutTSX) without a return value.
func resolveLayoutLogoutOps(layouts []stml.LayoutSpec, doc *openapi3.T, bearerAuth bool) {
	if !bearerAuth || doc == nil {
		return
	}
	logoutOp := stml_openapi.FindLogoutOp(doc)
	if logoutOp == "" {
		return
	}
	for i := range layouts {
		if layouts[i].Logout == nil {
			continue
		}
		if layouts[i].Logout.OperationID != "" {
			continue
		}
		layouts[i].Logout.OperationID = logoutOp
	}
}
