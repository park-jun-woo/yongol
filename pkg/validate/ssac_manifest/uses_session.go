//ff:func feature=validate type=util control=iteration dimension=2 topic=config-check
//ff:what usesSession — SSaC 에 @call session.* 호출이 있는지 확인

package ssac_manifest

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// usesSession reports whether any SSaC service func issues an @call whose
// model starts with "session." (e.g. session.GetUser, session.Put). This
// mirrors pkg/generate/prepared.sessionBackendFor detection so that
// validate and codegen agree on when the session subsystem is "in use".
func usesSession(fs *yongol.Fullstack) bool {
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type == "call" && strings.HasPrefix(seq.Model, "session.") {
				return true
			}
		}
	}
	return false
}
