//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what allPropsRequired — 스키마의 모든 top-level property가 required 목록에 있는지 판정

package stml_openapi

import "github.com/getkin/kin-openapi/openapi3"

// allPropsRequired reports whether every top-level property of a schema
// (allOf resolved one level) is required. Returns false when the schema has no
// properties.
func allPropsRequired(s *openapi3.Schema) bool {
	props := make(map[string]struct{})
	collectPropNames(props, s)
	if len(props) == 0 {
		return false
	}
	required := collectRequiredNames(s)
	for name := range props {
		if _, ok := required[name]; !ok {
			return false
		}
	}
	return true
}
