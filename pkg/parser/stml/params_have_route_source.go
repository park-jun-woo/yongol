//ff:func feature=stml-parse type=util control=iteration dimension=1
//ff:what ParamBind 슬라이스에 route. 접두사 소스가 있는지 확인
package stml

import "strings"

// paramsHaveRouteSource returns true if any param bind has a "route."
// prefixed source. Mirrors pkg/generate/react/has_route_source.go — kept
// in sync manually.
func paramsHaveRouteSource(params []ParamBind) bool {
	for _, param := range params {
		if strings.HasPrefix(param.Source, "route.") {
			return true
		}
	}
	return false
}
