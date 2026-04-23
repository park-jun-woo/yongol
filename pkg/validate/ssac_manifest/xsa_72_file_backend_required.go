//ff:func feature=validate type=rule control=sequence topic=config-check
//ff:what XSA-72 — @call file.* / storage.* requires manifest.file.backend

package ssac_manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xsa72FileBackendRequired validates XSA-72: if any SSaC service func
// calls file.* or storage.* built-ins, the manifest must declare
// file.backend. See XSA-70 for the same rationale applied to sessions.
func xsa72FileBackendRequired(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil {
		return nil
	}
	if !usesFile(fs) {
		return nil
	}
	if fs.Manifest != nil && fs.Manifest.File != nil && fs.Manifest.File.Backend != "" {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:    "manifest.yaml",
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: "[XSA-72] SSaC uses file.* / storage.* but manifest.file.backend is not declared",
		Advice:  "Declare manifest.file.backend (local | s3)",
	}}
}
