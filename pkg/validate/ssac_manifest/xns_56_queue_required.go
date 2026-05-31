//ff:func feature=validate type=rule control=sequence topic=config-check
//ff:what XNS-56 — @publish/@subscribe requires queue configuration in the manifest

package ssac_manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xns56QueueRequired validates XNS-56: if any SSaC ServiceFunc uses @publish
// or @subscribe, the manifest must define a queue backend (queue.backend).
func xns56QueueRequired(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil {
		return nil
	}
	if !usesQueue(fs) {
		return nil
	}
	g := fs.Ground()
	if g != nil && g.Config["queue.backend"] {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:    "manifest.yaml",
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: "[XNS-56] @publish/@subscribe is used but manifest.yaml queue.backend is not configured",
		Advice:  "Set manifest queue.backend (e.g. memory, kafka, or redis)",
	}}
}
