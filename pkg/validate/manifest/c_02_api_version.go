//ff:func feature=validate type=rule control=sequence topic=manifest-structural
//ff:what C-2 — validates that manifest apiVersion is "yongol/v1"

package manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// c02APIVersion validates that manifest.apiVersion is "yongol/v1".
func c02APIVersion(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.Manifest == nil {
		return nil
	}
	if fs.Manifest.APIVersion == "yongol/v1" {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:    "manifest.yaml",
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: "[C-2] apiVersion " + quoted(fs.Manifest.APIVersion) + " is not \"yongol/v1\"",
		Advice:  "Set apiVersion to yongol/v1",
	}}
}

