//ff:func feature=validate type=rule control=sequence topic=config-check
//ff:what XNS-48 — using currentUser requires claims to be enabled in the manifest

package ssac_manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xns48CurrentUserClaims validates XNS-48: when any SSaC sequence references
// currentUser (directly as an @auth or via currentUser.<field> inputs), the
// manifest must declare backend.auth.claims.
func xns48CurrentUserClaims(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil {
		return nil
	}
	if !usesCurrentUser(fs.ServiceFuncs) {
		return nil
	}
	g := fs.Ground()
	if g != nil && g.Config["backend.auth.claims"] {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:    "manifest.yaml",
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: "[XNS-48] currentUser is used but manifest.yaml backend.auth.claims is not defined",
		Advice:  "Enable manifest backend.auth.claims, or remove the use of currentUser",
	}}
}

