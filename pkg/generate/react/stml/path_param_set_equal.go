//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what 두 path-param 집합(paramName->type)이 같은 파라미터 이름 집합인지 비교한다 (비어 있으면 false)
package stml

// pathParamSetEqual reports whether two path-param maps cover the same
// non-empty set of parameter names. Only the names matter (the deleted
// item's GET is keyed by the same path params), so the types are ignored.
// Empty sets return false — a parameterless GET is never an item's own GET.
func pathParamSetEqual(a, b map[string]string) bool {
	if len(a) == 0 || len(a) != len(b) {
		return false
	}
	for k := range a {
		if _, ok := b[k]; !ok {
			return false
		}
	}
	return true
}
