//ff:func feature=validate type=util control=iteration dimension=1 topic=hurl-openapi
//ff:what descendAllOf — allOf 멤버를 순회하며 seg 프로퍼티를 찾는다

package hurl_openapi

import "github.com/getkin/kin-openapi/openapi3"

// descendAllOf inspects each allOf member for the given property name
// and returns the first match. Keeps descend() at nesting depth 2 by
// extracting the secondary iteration.
func descendAllOf(allOf openapi3.SchemaRefs, seg string) *openapi3.Schema {
	for _, ref := range allOf {
		if ref == nil || ref.Value == nil {
			continue
		}
		if prop, ok := ref.Value.Properties[seg]; ok && prop != nil {
			return prop.Value
		}
	}
	return nil
}
