//ff:func feature=gen-react type=util control=iteration dimension=1
//ff:what OpenAPI 경로에서 path parameter 이름 목록을 추출한다

package react

import "strings"

// extractPathParams extracts camelCase path parameter names from an OpenAPI path.
func extractPathParams(path string) []string {
	var params []string
	parts := strings.Split(path, "/")
	for _, p := range parts {
		if strings.HasPrefix(p, "{") && strings.HasSuffix(p, "}") {
			params = append(params, lcFirst(p[1:len(p)-1]))
		}
	}
	return params
}
