//ff:func feature=generate type=util control=sequence
//ff:what authFor — manifest.backend.auth 파생 (Mode 기본값 해석 단일화)

package prepared

import "github.com/park-jun-woo/yongol/pkg/yongol"

// authFor collapses the raw manifest.Auth into prepared.Auth with Mode
// already defaulted via manifest.Auth.ResolvedMode(). Every emitter
// that previously read a.Mode or a.ResolvedMode() now reads
// state.Auth.Mode, eliminating the BUG-009 class of inconsistency.
func authFor(fs *yongol.Fullstack) Auth {
	if fs == nil || fs.Manifest == nil || fs.Manifest.Backend.Auth == nil {
		return Auth{}
	}
	a := fs.Manifest.Backend.Auth
	return Auth{
		Present: true,
		Mode:    a.ResolvedMode(),
		Raw:     a,
	}
}
