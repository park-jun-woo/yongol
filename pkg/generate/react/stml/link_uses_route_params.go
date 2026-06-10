//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what LinkRef의 파라미터 소스에 route.* 가 있는지 판별한다 (useParams import 결정)
package stml

import (
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// linkUsesRouteParams reports whether any of the link's param sources is a
// route.<Name> reference, which requires the useParams() hook.
func linkUsesRouteParams(lr stmlparser.LinkRef) bool {
	for _, p := range lr.Params {
		if strings.HasPrefix(p.Source, "route.") {
			return true
		}
	}
	return false
}
