//ff:func feature=gen-react type=util control=iteration dimension=1
//ff:what ParamBind 슬라이스에서 route. 접두사 소스가 있는지 확인한다

package react

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// hasRouteSource returns true if any param bind has a "route." prefixed source.
func hasRouteSource(params []stml.ParamBind) bool {
	for _, param := range params {
		if strings.HasPrefix(param.Source, "route.") {
			return true
		}
	}
	return false
}
