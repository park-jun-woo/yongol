//ff:func feature=stml-parse type=util control=iteration dimension=1
//ff:what RoutePatternParams — 라우트 패턴의 :Name/:Name? 세그먼트 이름 추출 (optional 마커 ? strip)
package stml

import "strings"

// RoutePatternParams extracts the parameter segment names of a route
// pattern: every ":"-prefixed segment yields its name, with the optional
// marker ("?" suffix, react-router v6.5+) stripped. Names are returned
// exactly as declared — react-router's useParams() keys are
// case-sensitive, so consumers must compare case-exactly.
// e.g. "/unit-info/:UnitID/:PhotoID?" → ["UnitID", "PhotoID"].
func RoutePatternParams(pattern string) []string {
	if pattern == "" {
		return nil
	}
	var names []string
	for _, seg := range strings.Split(pattern, "/") {
		if !strings.HasPrefix(seg, ":") {
			continue
		}
		name := strings.TrimSuffix(strings.TrimPrefix(seg, ":"), "?")
		if name == "" {
			continue
		}
		names = append(names, name)
	}
	return names
}
