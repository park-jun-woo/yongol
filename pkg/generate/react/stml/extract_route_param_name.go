//ff:func feature=stml-gen type=util control=sequence
//ff:what ParamBind에서 route. 접두사가 있으면 파라미터 이름을 반환한다
package stml

import (
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func extractRouteParamName(p stmlparser.ParamBind) string {
	if strings.HasPrefix(p.Source, "route.") {
		return strings.TrimPrefix(p.Source, "route.")
	}
	return ""
}
