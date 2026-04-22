//ff:func feature=validate type=util control=iteration dimension=1 topic=manifest-cors
//ff:what corsAllowOriginsHasWildcard — allow_origins 목록에 "*" 포함 여부

package manifest

// corsAllowOriginsHasWildcard reports whether the allow_origins list contains
// the wildcard "*".
func corsAllowOriginsHasWildcard(origins []string) bool {
	for _, o := range origins {
		if o == "*" {
			return true
		}
	}
	return false
}
