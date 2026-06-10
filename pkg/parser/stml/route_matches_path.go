//ff:func feature=stml-parse type=util control=iteration dimension=1
//ff:what RouteMatchesPath — 정적 경로가 라우트 패턴(:param 세그먼트 허용)에 매칭되는지 판정
package stml

import "strings"

// RouteMatchesPath reports whether a concrete static path (no params)
// matches a route pattern: segments are compared one-to-one, and a
// ":"-prefixed pattern segment matches any single non-empty path segment.
func RouteMatchesPath(pattern, path string) bool {
	if pattern == path {
		return true
	}
	pSegs := strings.Split(strings.Trim(pattern, "/"), "/")
	aSegs := strings.Split(strings.Trim(path, "/"), "/")
	if len(pSegs) != len(aSegs) {
		return false
	}
	for i := range pSegs {
		if strings.HasPrefix(pSegs[i], ":") && aSegs[i] != "" {
			continue
		}
		if pSegs[i] != aSegs[i] {
			return false
		}
	}
	return true
}
