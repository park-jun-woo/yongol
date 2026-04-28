//ff:func feature=validate type=util control=sequence topic=hurl-openapi
//ff:what descend — schema 를 한 세그먼트만큼 내려간다 (array items / properties / allOf)

package hurl_openapi

import (
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// descend advances one segment through a schema. Array segments pass
// through to `items`; object segments look up `properties`; allOf
// members are inspected in order for the property.
func descend(s *openapi3.Schema, seg string) *openapi3.Schema {
	if s == nil {
		return nil
	}
	if strings.HasPrefix(seg, "[") && strings.HasSuffix(seg, "]") {
		if s.Items != nil && s.Items.Value != nil {
			return s.Items.Value
		}
		return nil
	}
	if prop, ok := s.Properties[seg]; ok && prop != nil {
		return prop.Value
	}
	return descendAllOf(s.AllOf, seg)
}
