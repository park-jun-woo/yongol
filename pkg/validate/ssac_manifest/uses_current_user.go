//ff:func feature=validate type=util control=iteration dimension=3 topic=config-check
//ff:what usesCurrentUser — SSaC 시퀀스가 currentUser/@auth 를 사용하는지 확인

package ssac_manifest

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// usesCurrentUser reports whether any SSaC sequence uses currentUser or @auth.
func usesCurrentUser(funcs []ssac.ServiceFunc) bool {
	for _, fn := range funcs {
		for _, seq := range fn.Sequences {
			if seq.Type == "auth" {
				return true
			}
			for _, v := range seq.Inputs {
				if strings.HasPrefix(v, "currentUser.") {
					return true
				}
			}
		}
	}
	return false
}
