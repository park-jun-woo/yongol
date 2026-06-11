//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what firstRequiredSegment — 라우트 패턴의 첫 필수(:Name, ? 없음) 세그먼트 반환 (없으면 "")

package stml_openapi

import "strings"

// firstRequiredSegment returns the first required (":Name" without the "?"
// optional suffix) segment of a resolved route pattern, "" when the route
// has none — the TM-34 redirect-target judgment shared by TM-42.
func firstRequiredSegment(pattern string) string {
	for _, seg := range strings.Split(pattern, "/") {
		if strings.HasPrefix(seg, ":") && !strings.HasSuffix(seg, "?") {
			return seg
		}
	}
	return ""
}
