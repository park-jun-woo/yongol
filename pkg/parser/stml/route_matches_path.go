//ff:func feature=stml-parse type=util control=iteration dimension=1
//ff:what RouteMatchesPath — 정적 경로가 라우트 패턴(:param 세그먼트, 말미 :param? 생략 허용)에 매칭되는지 판정
package stml

import "strings"

// RouteMatchesPath reports whether a concrete static path (no params)
// matches a route pattern: segments are compared one-to-one, a
// ":"-prefixed pattern segment matches any single non-empty path segment,
// and trailing optional segments (":Name?", react-router v6.5+) may be
// omitted from the path.
func RouteMatchesPath(pattern, path string) bool {
	if pattern == path {
		return true
	}
	pSegs := strings.Split(strings.Trim(pattern, "/"), "/")
	aSegs := strings.Split(strings.Trim(path, "/"), "/")
	if len(aSegs) > len(pSegs) {
		return false
	}
	for i, aSeg := range aSegs {
		if strings.HasPrefix(pSegs[i], ":") && aSeg != "" {
			continue
		}
		if pSegs[i] != aSeg {
			return false
		}
	}
	// surplus pattern segments must all be trailing optional (":Name?")
	for _, pSeg := range pSegs[len(aSegs):] {
		if !strings.HasPrefix(pSeg, ":") || !strings.HasSuffix(pSeg, "?") {
			return false
		}
	}
	return true
}
