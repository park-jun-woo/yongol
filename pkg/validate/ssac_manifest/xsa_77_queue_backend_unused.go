//ff:func feature=validate type=rule control=sequence topic=config-check
//ff:what XSA-77 — manifest.queue.backend declared but SSaC never @publish/@subscribe

package ssac_manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xsa77QueueBackendUnused validates XSA-77: the manifest declares a queue
// backend but no SSaC service func declares @publish or @subscribe. See
// XSA-74 for the same rationale applied to sessions.
func xsa77QueueBackendUnused(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil {
		return nil
	}
	if fs.Manifest == nil || fs.Manifest.Queue == nil || fs.Manifest.Queue.Backend == "" {
		return nil
	}
	if usesQueue(fs) {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:    "manifest.yaml",
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelWarning,
		Message: "[XSA-77] manifest.queue.backend is declared but no SSaC function uses @publish / @subscribe",
		Advice:  "Remove manifest.queue.backend or add an @publish / @subscribe sequence",
	}}
}
