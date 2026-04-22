//ff:func feature=validate type=rule control=sequence topic=manifest-structural
//ff:what C-5 — validates that manifest backend.module is non-empty

package manifest

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// c05BackendModule validates that manifest.backend.module is non-empty.
func c05BackendModule(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.Manifest == nil {
		return nil
	}
	if strings.TrimSpace(fs.Manifest.Backend.Module) != "" {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:    "manifest.yaml",
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: "[C-5] backend.module is empty",
		Advice:  "Set backend.module to the Go module path (e.g. github.com/org/project)",
	}}
}
