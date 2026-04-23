//ff:func feature=validate type=util control=iteration dimension=2 topic=config-check
//ff:what usesCache — SSaC 에 @call cache.* 호출이 있는지 확인

package ssac_manifest

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// usesCache reports whether any SSaC service func issues an @call whose
// model starts with "cache." (e.g. cache.Get, cache.Set). Mirrors
// pkg/generate/gogin/boot.hasCache detection so validate and codegen
// agree on when the cache subsystem is "in use".
func usesCache(fs *yongol.Fullstack) bool {
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type == "call" && strings.HasPrefix(seq.Model, "cache.") {
				return true
			}
		}
	}
	return false
}
