//ff:func feature=validate type=util control=iteration dimension=1 topic=scenario-check
//ff:what normalizeHurlPath — Hurl URL 경로를 정규화 세그먼트로 변환

package openapi_hurl

import (
	"regexp"
	"strings"
)

var (
	reHurlVar     = regexp.MustCompile(`\{\{.+?\}\}`)
	reHurlNumeric = regexp.MustCompile(`^\d+$`)
)

// normalizeHurlPath converts a Hurl URL path to normalized segments.
// {{variable}} -> ":param", pure numeric literals -> ":param"
func normalizeHurlPath(path string) []string {
	path = strings.TrimSpace(path)
	if idx := strings.Index(path, "?"); idx >= 0 {
		path = path[:idx]
	}
	parts := strings.Split(strings.Trim(path, "/"), "/")
	var segs []string
	for _, p := range parts {
		if p == "" {
			continue
		}
		if reHurlVar.MatchString(p) || reHurlNumeric.MatchString(p) {
			segs = append(segs, ":param")
		} else {
			segs = append(segs, p)
		}
	}
	return segs
}
