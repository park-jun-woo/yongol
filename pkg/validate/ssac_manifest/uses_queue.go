//ff:func feature=validate type=util control=iteration dimension=2 topic=config-check
//ff:what usesQueue — SSaC에 @subscribe 또는 publish 시퀀스가 있는지 확인

package ssac_manifest

import "github.com/park-jun-woo/yongol/pkg/yongol"

// usesQueue reports whether any SSaC function declares @subscribe or any
// sequence is a publish step.
func usesQueue(fs *yongol.Fullstack) bool {
	for _, fn := range fs.ServiceFuncs {
		if fn.Subscribe != nil {
			return true
		}
		for _, seq := range fn.Sequences {
			if seq.Type == "publish" {
				return true
			}
		}
	}
	return false
}
