//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what ParamBind 슬라이스에서 route. 접두사를 가진 고유 파라미터 이름을 추출한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

func extractRouteParamNames(params []stmlparser.ParamBind) []string {
	var routeParams []string
	seen := map[string]bool{}
	for _, p := range params {
		name := extractRouteParamName(p)
		if name != "" && !seen[name] {
			routeParams = append(routeParams, name)
			seen[name] = true
		}
	}
	return routeParams
}
