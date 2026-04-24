//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what isLoginShape — SSaC ServiceFunc 가 @verify-password 시퀀스를 포함하는지 판정

package hurl

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// isLoginShape returns true when the service func contains at least one
// @verify-password sequence. That DSL token is yongol's timing-safe
// login primitive (see ssac/parse_verify_password.go) and is not used
// anywhere else — presence of any SeqVerifyPassword sequence is a
// sufficient signal that the op is a login endpoint, regardless of
// operationId naming.
func isLoginShape(fn *ssac.ServiceFunc) bool {
	if fn == nil {
		return false
	}
	for _, seq := range fn.Sequences {
		if seq.Type == ssac.SeqVerifyPassword {
			return true
		}
	}
	return false
}
