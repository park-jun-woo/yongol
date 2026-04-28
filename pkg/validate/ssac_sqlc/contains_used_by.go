//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-sqlc
//ff:what containsUsedBy — used_by 슬라이스에 method 가 포함되어 있는지 판정

package ssac_sqlc

// containsUsedBy reports whether usedBy contains method.
func containsUsedBy(usedBy []string, method string) bool {
	for _, m := range usedBy {
		if m == method {
			return true
		}
	}
	return false
}
