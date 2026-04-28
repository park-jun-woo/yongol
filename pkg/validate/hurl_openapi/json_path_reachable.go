//ff:func feature=validate type=util control=iteration dimension=1 topic=hurl-openapi
//ff:what jsonPathReachable — 점 표기 JSONPath 가 OpenAPI schema 에서 도달 가능한지 판정

package hurl_openapi

import (
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// jsonPathReachable walks a dotted JSONPath (`$.user.id`,
// `$.items[0].name`) through an OpenAPI schema and returns true when
// every segment resolves. Array indexes (`[n]`) descend into the
// `items` schema; unknown segments short-circuit to false.
//
// Wildcards (`$..email`, `$[*]`) are conservatively treated as
// reachable — hurl users often lean on them for flexible assertions
// and a strict walker would produce noisy false positives.
func jsonPathReachable(path string, schema *openapi3.Schema) bool {
	if schema == nil || path == "" {
		return false
	}
	if strings.Contains(path, "..") || strings.Contains(path, "[*]") || strings.Contains(path, "[?") {
		return true
	}
	segs := parseJSONPath(path)
	cur := schema
	for _, seg := range segs {
		cur = descend(cur, seg)
		if cur == nil {
			return false
		}
	}
	return true
}
