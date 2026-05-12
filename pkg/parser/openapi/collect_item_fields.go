//ff:func feature=openapi-parse type=parser control=iteration dimension=1
//ff:what collectItemFields — 단일 배열 프로퍼티의 items 스키마에서 필드 이름 집합 반환

package openapi

import "github.com/getkin/kin-openapi/openapi3"

// collectItemFields returns the property name set of an array schema's
// items, or nil if the property is not an array or has no items.
func collectItemFields(propRef *openapi3.SchemaRef) map[string]bool {
	if propRef.Value == nil || propRef.Value.Type == nil || !propRef.Value.Type.Is("array") {
		return nil
	}
	items := propRef.Value.Items
	if items == nil || items.Value == nil {
		return nil
	}
	itemFields := make(map[string]bool)
	for itemPropName := range items.Value.Properties {
		itemFields[itemPropName] = true
	}
	return itemFields
}
