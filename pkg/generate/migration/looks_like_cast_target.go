//ff:func feature=migration type=util control=iteration dimension=1
//ff:what looksLikeCastTarget — ::<target> 의 target 이 단순 식별자/ "quoted" 인지 판정
package migration

import "strings"

// looksLikeCastTarget reports whether s is a simple identifier such as
// `text`, `"MyType"`, or `integer[]`.
func looksLikeCastTarget(s string) bool {
	if s == "" {
		return false
	}
	for strings.HasSuffix(s, "[]") {
		s = strings.TrimSpace(strings.TrimSuffix(s, "[]"))
	}
	if strings.HasPrefix(s, `"`) && strings.HasSuffix(s, `"`) {
		return true
	}
	for _, r := range s {
		if !isIdentOrSpaceRune(r) {
			return false
		}
	}
	return true
}
