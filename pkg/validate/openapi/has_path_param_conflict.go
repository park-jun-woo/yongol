//ff:func feature=validate type=util control=iteration dimension=1 topic=openapi-structural
//ff:what hasPathParamConflict — path 내 중복 {param} 감지 헬퍼

package openapi

import "strings"

func hasPathParamConflict(path string) bool {
	seen := map[string]bool{}
	for _, seg := range strings.Split(path, "/") {
		if !strings.HasPrefix(seg, "{") || !strings.HasSuffix(seg, "}") {
			continue
		}
		name := seg[1 : len(seg)-1]
		if seen[name] {
			return true
		}
		seen[name] = true
	}
	return false
}
