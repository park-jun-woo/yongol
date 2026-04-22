//ff:func feature=validate type=util control=iteration dimension=1 topic=openapi-structural
//ff:what collectDeclaredPathParams — path-level + operation-level parameters 의 in:path name 집합 추출

package openapi

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// collectDeclaredPathParams returns the set of parameter names declared with
// `in: path` across the given path-level and operation-level parameter lists.
// OpenAPI 3.x allows parameters on either level; unions both.
func collectDeclaredPathParams(pathParams, opParams openapi3.Parameters) map[string]bool {
	set := map[string]bool{}
	all := append(openapi3.Parameters{}, pathParams...)
	all = append(all, opParams...)
	for _, ref := range all {
		if ref == nil || ref.Value == nil {
			continue
		}
		if ref.Value.In == "path" && ref.Value.Name != "" {
			set[ref.Value.Name] = true
		}
	}
	return set
}
