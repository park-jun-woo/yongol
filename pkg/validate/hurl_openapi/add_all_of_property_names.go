//ff:func feature=validate type=util control=iteration dimension=1 topic=hurl-openapi
//ff:what addAllOfPropertyNames — allOf 멤버의 property 이름을 out 세트에 추가

package hurl_openapi

import "github.com/getkin/kin-openapi/openapi3"

// addAllOfPropertyNames merges each allOf member's top-level properties
// into out. Kept as a dedicated helper so schemaPropertyNames stays at
// depth 2 (sequence-of-for, not for-within-for).
func addAllOfPropertyNames(out map[string]struct{}, allOf openapi3.SchemaRefs) {
	for _, ref := range allOf {
		if ref == nil || ref.Value == nil {
			continue
		}
		for name := range ref.Value.Properties {
			out[name] = struct{}{}
		}
	}
}
