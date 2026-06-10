//ff:func feature=validate type=util control=sequence topic=stml-openapi
//ff:what itemParamField — item.<Field> 소스에서 최상위 필드명을 추출 (item.* 아니면 false)

package stml_openapi

import "strings"

// itemParamField returns the top-level field name of an item.<Field> param
// source ("item.photo.id" → "photo") and true, or "" and false when the
// source is not an item.* reference.
func itemParamField(source string) (string, bool) {
	if !strings.HasPrefix(source, "item.") {
		return "", false
	}
	field := strings.TrimPrefix(source, "item.")
	if i := strings.IndexByte(field, '.'); i >= 0 {
		field = field[:i]
	}
	return field, true
}
