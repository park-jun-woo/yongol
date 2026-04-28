//ff:func feature=validate type=util control=iteration dimension=1 topic=hurl-openapi
//ff:what normalizeHurlPath — Hurl 요청 path 를 세그먼트 배열로 정규화

package hurl_openapi

import (
	"regexp"
	"strings"
)

var (
	reHurlVar     = regexp.MustCompile(`\{\{.+?\}\}`)
	reHurlNumeric = regexp.MustCompile(`^\d+$`)
	reOpenAPIVar  = regexp.MustCompile(`^\{.+\}$`)
)

// normalizeHurlPath converts a Hurl request path into segments suitable
// for comparison with normalizeOpenAPIPath output. `{{var}}` and pure
// numeric literals collapse to `:param` so path-parameter positions
// match whatever OpenAPI declared.
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
