//ff:func feature=ssacmeta type=util control=selection
//ff:what truthy — Go 기본 타입 v 가 truthy 한지 (zero value 가 아닌지) 판정

package ssacmeta

// truthy reports whether v is "truthy" in the when: DSL sense — i.e. a
// non-zero value for its underlying type.
func truthy(v any) bool {
	switch x := v.(type) {
	case nil:
		return false
	case bool:
		return x
	case string:
		return x != ""
	case int:
		return x != 0
	case int64:
		return x != 0
	case float64:
		return x != 0
	default:
		return true
	}
}
