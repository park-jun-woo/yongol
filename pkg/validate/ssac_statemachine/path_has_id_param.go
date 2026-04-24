//ff:func feature=validate type=util control=iteration dimension=1 topic=states
//ff:what pathHasIDParam — OpenAPI path 템플릿에 `{id}` 파라미터 포함 여부

package ssac_statemachine

import (
	"strings"
)

// pathHasIDParam reports whether an OpenAPI path template contains an {id}
// parameter. Lenient on case so `{ID}` / `{Id}` pass too.
func pathHasIDParam(path string) bool {
	for _, part := range strings.Split(path, "/") {
		if !strings.HasPrefix(part, "{") || !strings.HasSuffix(part, "}") {
			continue
		}
		name := strings.TrimSuffix(strings.TrimPrefix(part, "{"), "}")
		if strings.EqualFold(name, "id") {
			return true
		}
	}
	return false
}
