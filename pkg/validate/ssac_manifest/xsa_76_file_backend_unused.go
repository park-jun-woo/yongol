//ff:func feature=validate type=rule control=sequence topic=config-check
//ff:what XSA-76 — manifest.file.backend declared but SSaC never calls file.*/storage.*

package ssac_manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xsa76FileBackendUnused validates XSA-76: the manifest declares a file
// backend but no SSaC service func calls file.* or storage.*. See XSA-74
// for the same rationale applied to sessions.
func xsa76FileBackendUnused(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil {
		return nil
	}
	if fs.Manifest == nil || fs.Manifest.File == nil || fs.Manifest.File.Backend == "" {
		return nil
	}
	if usesFile(fs) {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:    "manifest.yaml",
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelWarning,
		Message: "[XSA-76] manifest.file.backend is declared but no SSaC function uses file.* / storage.*",
		Advice:  "Remove manifest.file.backend or add an @call file.* / storage.* sequence",
	}}
}
