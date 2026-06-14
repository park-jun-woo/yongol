//ff:func feature=stml-gen type=util control=sequence
//ff:what ParamBind Source가 route.<Name>이면 세그먼트 이름을 반환한다 (BUG-136)
package stml

import "strings"

// routeSegmentName returns the route segment name for a "route.<Name>" source
// (e.g. "route.BuildingID" → "BuildingID", true). Non-route or empty sources
// return ("", false).
func routeSegmentName(source string) (string, bool) {
	if !strings.HasPrefix(source, "route.") {
		return "", false
	}
	name := strings.TrimPrefix(source, "route.")
	if name == "" {
		return "", false
	}
	return name, true
}
