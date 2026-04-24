//ff:func feature=gen-hurl type=test-helper control=sequence
//ff:what loginServiceFunc — login-shape SSaC ServiceFunc 픽스처 생성

package hurl

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// loginServiceFunc returns a minimal login-shape SSaC ServiceFunc
// fixture — one @verify-password sequence is sufficient.
func loginServiceFunc(opID string) ssac.ServiceFunc {
	return ssac.ServiceFunc{
		Name: opID,
		Sequences: []ssac.Sequence{
			{
				Type:         ssac.SeqVerifyPassword,
				Model:        "User",
				EmailCol:     "email",
				EmailExpr:    "request.email",
				HashCol:      "password_hash",
				PasswordExpr: "request.password",
			},
		},
	}
}
