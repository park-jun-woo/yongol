//ff:func feature=gen-react type=util control=iteration dimension=1
//ff:what resolveLogoutRefreshArgs -- logout op의 requestBody에 refresh 프로퍼티가 있으면 LogoutSpec.RefreshBodyKey를 채운다

package react

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func resolveLogoutRefreshArgs(layouts []stml.LayoutSpec, doc *openapi3.T, refreshField string) {
	if refreshField == "" || doc == nil || doc.Paths == nil {
		return
	}
	for i := range layouts {
		lo := layouts[i].Logout
		if lo == nil || lo.OperationID == "" {
			continue
		}
		op := findOpByID(doc, lo.OperationID)
		if op == nil {
			continue
		}
		lo.RefreshBodyKey = refreshBodyKey(op, refreshField)
	}
}
