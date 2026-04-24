//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what hasUserCreatePost — @post <Model>.Create 에 PasswordHash 인자가 있는지 확인

package hurl

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// hasUserCreatePost checks for a companion `@post <Model>.Create` sequence
// that wires a PasswordHash-like column. Used only to emit a WARNING
// when signup shape is suspicious (HashPassword without a Create post).
func hasUserCreatePost(fn *ssac.ServiceFunc) bool {
	if fn == nil {
		return false
	}
	for _, seq := range fn.Sequences {
		if isUserCreateWithPasswordHash(seq) {
			return true
		}
	}
	return false
}
