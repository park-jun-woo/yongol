//ff:func feature=stml-parse type=util control=iteration dimension=1
//ff:what appendRouteParams — ParamBind의 route.<Name> 소스를 첫 등장 순서로 중복 없이 누적
package stml

import "strings"

// appendRouteParams appends the route.<Name> sources found in binds to
// params, skipping names already in seen — the first appearance (and its
// required flag) wins.
func appendRouteParams(params []routeParam, seen map[string]bool, binds []ParamBind, required bool) []routeParam {
	for _, b := range binds {
		if !strings.HasPrefix(b.Source, "route.") {
			continue
		}
		name := strings.TrimPrefix(b.Source, "route.")
		if name == "" || seen[name] {
			continue
		}
		seen[name] = true
		params = append(params, routeParam{Name: name, Required: required})
	}
	return params
}
