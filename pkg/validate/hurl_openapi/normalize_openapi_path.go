//ff:func feature=validate type=util control=iteration dimension=1 topic=hurl-openapi
//ff:what normalizeOpenAPIPath — OpenAPI path 를 세그먼트 배열로 정규화

package hurl_openapi

import "strings"

// normalizeOpenAPIPath converts a path declaration like
// `/users/{id}/posts/{postId}` into `["users", ":param", "posts",
// ":param"]`.
func normalizeOpenAPIPath(path string) []string {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	var segs []string
	for _, p := range parts {
		if p == "" {
			continue
		}
		if reOpenAPIVar.MatchString(p) {
			segs = append(segs, ":param")
		} else {
			segs = append(segs, p)
		}
	}
	return segs
}
