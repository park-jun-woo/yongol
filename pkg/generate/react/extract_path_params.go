//ff:func feature=gen-react type=util control=iteration dimension=1
//ff:what OpenAPI 경로에서 path parameter 이름 목록을 원본 그대로 추출한다

package react

import "strings"

// extractPathParams extracts path parameter names from an OpenAPI path,
// preserving the original template name verbatim (no casing transform).
//
// The template name ({contractId}), the openapi-typescript Req<K> type, and
// the call-site argument key all derive from the same OpenAPI parameter name,
// so the extraction key must match it exactly. Applying a Go-identifier casing
// transform (e.g. ToGoCamel promoting "Id" -> "ID") would diverge from those
// keys and break runtime path templating (BUG-109).
func extractPathParams(path string) []string {
	var params []string
	parts := strings.Split(path, "/")
	for _, p := range parts {
		if strings.HasPrefix(p, "{") && strings.HasSuffix(p, "}") {
			params = append(params, p[1:len(p)-1])
		}
	}
	return params
}
