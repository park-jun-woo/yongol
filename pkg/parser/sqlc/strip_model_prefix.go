//ff:func feature=orchestrator type=util control=sequence dimension=1
//ff:what stripModelPrefix — query 이름에서 모델 prefix를 제거 (없으면 원본)
package sqlc

import "strings"

// stripModelPrefix removes the model PascalCase prefix from the query name.
// If the name does not start with the model prefix, the original name is
// returned unchanged.
func stripModelPrefix(name, model string) string {
	if model == "" {
		return name
	}
	if !strings.HasPrefix(name, model) {
		return name
	}
	stripped := strings.TrimPrefix(name, model)
	if stripped == "" {
		return name
	}
	return stripped
}
