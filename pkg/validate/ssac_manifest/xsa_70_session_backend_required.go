//ff:func feature=validate type=rule control=sequence topic=config-check
//ff:what XSA-70 — @call session.* requires manifest.session.backend

package ssac_manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xsa70SessionBackendRequired validates XSA-70: if any SSaC service func
// calls session.* built-ins, the manifest must declare session.backend.
// Missing declaration is an ERROR because yongol codegen dereferences
// fs.Manifest.Session.Backend unconditionally when SSaC uses session —
// letting it through validate produces a nil-pointer panic at generate
// time.
func xsa70SessionBackendRequired(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil {
		return nil
	}
	if !usesSession(fs) {
		return nil
	}
	if fs.Manifest != nil && fs.Manifest.Session != nil && fs.Manifest.Session.Backend != "" {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:    "manifest.yaml",
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: "[XSA-70] SSaC uses session.* but manifest.session.backend is not declared",
		Advice:  "Declare manifest.session.backend (memory | postgres)",
	}}
}
