//ff:func feature=manifest type=parser control=sequence
//ff:what mapValueType — object 의 additionalProperties 에서 맵 값 타입(또는 자유형 마커 "*") 추출
package openapi

import "github.com/getkin/kin-openapi/openapi3"

// mapValueType extracts the value type of an object(map) from its
// additionalProperties. additionalProperties: { type: <value> } → "<value>".
// additionalProperties: true (Has) or unspecified → free-form marker "*".
func mapValueType(ap openapi3.AdditionalProperties) string {
	if ap.Schema == nil || ap.Schema.Value == nil {
		return "*"
	}
	if valTypes := ap.Schema.Value.Type.Slice(); len(valTypes) > 0 {
		return valTypes[0]
	}
	return "*"
}
