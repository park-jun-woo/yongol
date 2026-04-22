//ff:func feature=validate type=util control=sequence topic=func-check
//ff:what toCamelKey — PascalCase → camelCase 변환 (Func 어노테이션 키 정규화)

package ssac_func

// toCamelKey converts PascalCase (@call style "HashPassword") to the
// camelCase @func annotation form ("hashPassword") used as Ground.Types
// and Ground.Schemas key for Func.request lookups.
func toCamelKey(pascal string) string {
	if pascal == "" {
		return ""
	}
	first := pascal[0]
	if first >= 'A' && first <= 'Z' {
		return string(first+('a'-'A')) + pascal[1:]
	}
	return pascal
}
