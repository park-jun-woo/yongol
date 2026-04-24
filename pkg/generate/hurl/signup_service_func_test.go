//ff:func feature=gen-hurl type=test-helper control=sequence
//ff:what signupServiceFunc — signup-shape SSaC ServiceFunc 픽스처 생성

package hurl

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// signupServiceFunc returns a minimal signup-shape SSaC ServiceFunc
// fixture. The HashPassword @call + User.Create @post pair is what
// detectAuthOps matches on.
func signupServiceFunc(opID string) ssac.ServiceFunc {
	return ssac.ServiceFunc{
		Name: opID,
		Sequences: []ssac.Sequence{
			{
				Type:  ssac.SeqCall,
				Model: "auth.HashPassword",
				Inputs: map[string]string{
					"Password": "request.password",
				},
			},
			{
				Type:  ssac.SeqPost,
				Model: "User.Create",
				Inputs: map[string]string{
					"Email":        "request.email",
					"PasswordHash": "hp.HashedPassword",
				},
			},
		},
	}
}
