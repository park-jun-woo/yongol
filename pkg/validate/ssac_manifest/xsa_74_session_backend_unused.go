//ff:func feature=validate type=rule control=sequence topic=config-check
//ff:what XSA-74 — manifest.session.backend declared but SSaC never calls session.*

package ssac_manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xsa74SessionBackendUnused validates XSA-74: the manifest declares a
// session backend but no SSaC service func calls session.*. yongol will
// still emit the session init block, so the runtime wires unused
// infrastructure — flagged as WARNING for hygiene.
func xsa74SessionBackendUnused(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil {
		return nil
	}
	if fs.Manifest == nil || fs.Manifest.Session == nil || fs.Manifest.Session.Backend == "" {
		return nil
	}
	if usesSession(fs) {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:    "manifest.yaml",
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelWarning,
		Message: "[XSA-74] manifest.session.backend is declared but no SSaC function uses session.*",
		Advice:  "Remove manifest.session.backend or add an @call session.* sequence",
	}}
}
