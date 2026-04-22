//ff:func feature=validate type=util control=iteration dimension=1 topic=scenario-check
//ff:what normalizeOpenAPIPath — OpenAPI 경로를 정규화 세그먼트로 변환

package openapi_hurl

import (
	"regexp"
	"strings"
)

var reOpenAPIParam = regexp.MustCompile(`^\{.+\}$`)

// normalizeOpenAPIPath converts an OpenAPI path to normalized segments.
// {param} -> ":param"
func normalizeOpenAPIPath(path string) []string {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	var segs []string
	for _, p := range parts {
		if p == "" {
			continue
		}
		if reOpenAPIParam.MatchString(p) {
			segs = append(segs, ":param")
		} else {
			segs = append(segs, p)
		}
	}
	return segs
}
