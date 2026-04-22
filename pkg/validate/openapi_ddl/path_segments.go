//ff:func feature=validate type=util control=iteration dimension=1 topic=openapi-ddl
//ff:what pathSegments — OpenAPI path에서 concrete segment만 역순으로 반환

package openapi_ddl

import "strings"

// pathSegments returns concrete (non-{param}) segments from an OpenAPI path.
func pathSegments(p string) []string {
	parts := strings.Split(p, "/")
	out := make([]string, 0, len(parts))
	for _, s := range parts {
		if s == "" || strings.HasPrefix(s, "{") {
			continue
		}
		out = append(out, s)
	}
	// reverse so caller checks the last segment first (most specific)
	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}
	return out
}
