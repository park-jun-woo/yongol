//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what data-route 패턴 문자열에서 :paramName 을 추출한다 (optional 마커 ? strip)
package stml

import "strings"

// extractRoutePatternParams extracts parameter names from a route pattern.
// Optional segment markers (":Name?", react-router v6.5+) are stripped so
// useParams merging gets the exact param name.
// e.g. "/buildings/:buildingId/units/:id" → ["buildingId", "id"]
// and "/unit-info/:UnitID/:PhotoID?" → ["UnitID", "PhotoID"].
func extractRoutePatternParams(route string) []string {
	if route == "" {
		return nil
	}
	var params []string
	for _, seg := range strings.Split(route, "/") {
		if strings.HasPrefix(seg, ":") {
			params = append(params, strings.TrimSuffix(strings.TrimPrefix(seg, ":"), "?"))
		}
	}
	return params
}
