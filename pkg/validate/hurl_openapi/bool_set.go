//ff:func feature=validate type=util control=iteration dimension=1 topic=hurl-openapi
//ff:what boolSet — struct{} set 을 bool map 으로 변환 (sortedKeys 재사용)

package hurl_openapi

// boolSet converts a struct{} set into a bool map so sortedKeys can be
// reused. A dedicated helper keeps the two callers symmetric.
func boolSet(s map[string]struct{}) map[string]bool {
	out := make(map[string]bool, len(s))
	for k := range s {
		out[k] = true
	}
	return out
}
