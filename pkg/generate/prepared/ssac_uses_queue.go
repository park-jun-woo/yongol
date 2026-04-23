//ff:func feature=generate type=util control=iteration dimension=2
//ff:what ssacUsesQueue — SSaC @subscribe 또는 publish 시퀀스 존재 여부

package prepared

import "github.com/park-jun-woo/yongol/pkg/yongol"

// ssacUsesQueue reports whether any SSaC function declares @subscribe or
// any sequence is a publish step. Mirrors
// pkg/validate/ssac_manifest.usesQueue.
func ssacUsesQueue(fs *yongol.Fullstack) bool {
	if fs == nil {
		return false
	}
	for _, fn := range fs.ServiceFuncs {
		if funcUsesQueue(fn) {
			return true
		}
	}
	return false
}
