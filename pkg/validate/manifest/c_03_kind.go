//ff:func feature=validate type=rule control=sequence topic=manifest-structural
//ff:what C-3 — validates that manifest kind is "Project"

package manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// c03Kind validates that manifest.kind is "Project".
func c03Kind(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.Manifest == nil {
		return nil
	}
	if fs.Manifest.Kind == "Project" {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:    "manifest.yaml",
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: "[C-3] kind " + quoted(fs.Manifest.Kind) + " is not \"Project\"",
		Advice:  "Set kind to Project",
	}}
}
