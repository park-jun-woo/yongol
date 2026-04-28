//ff:func feature=validate type=util control=iteration dimension=1 topic=hurl-openapi
//ff:what parseJSONPath — `$.a.b[0].c` 를 세그먼트 배열로 분해

package hurl_openapi

import "strings"

// parseJSONPath splits `$.a.b[0].c` into `["a", "b", "[0]", "c"]`.
func parseJSONPath(path string) []string {
	p := strings.TrimPrefix(path, "$")
	p = strings.TrimPrefix(p, ".")
	var out []string
	var cur strings.Builder
	for i := 0; i < len(p); i++ {
		i = parseJSONPathChar(p, i, &cur, &out)
	}
	if cur.Len() > 0 {
		out = append(out, cur.String())
	}
	return out
}
