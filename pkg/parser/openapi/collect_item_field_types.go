//ff:func feature=openapi-parse type=parser control=iteration dimension=1
//ff:what collectItemFieldTypes — 단일 배열 프로퍼티의 items 스키마에서 필드명→타입 맵 반환

package openapi

import "github.com/getkin/kin-openapi/openapi3"

// collectItemFieldTypes returns the property name → OpenAPI type map of an
// array schema's items (e.g. {"id": "integer", "caption": "string"}), or nil
// if the property is not an array or has no items. Properties without an
// explicit type are skipped.
func collectItemFieldTypes(propRef *openapi3.SchemaRef) map[string]string {
	if propRef.Value == nil || propRef.Value.Type == nil || !propRef.Value.Type.Is("array") {
		return nil
	}
	items := propRef.Value.Items
	if items == nil || items.Value == nil {
		return nil
	}
	itemTypes := make(map[string]string)
	for itemPropName, itemPropRef := range items.Value.Properties {
		if itemPropRef == nil || itemPropRef.Value == nil || itemPropRef.Value.Type == nil {
			continue
		}
		types := itemPropRef.Value.Type.Slice()
		if len(types) == 0 {
			continue
		}
		itemTypes[itemPropName] = types[0]
	}
	if len(itemTypes) == 0 {
		return nil
	}
	return itemTypes
}
