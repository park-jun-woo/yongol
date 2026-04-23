//ff:func feature=gen-gogin type=util control=sequence topic=observability
//ff:what hasOtel — manifest.backend.observability.tracing.enabled 여부 (기본 false)

package boot

import "github.com/park-jun-woo/yongol/pkg/yongol"

// hasOtel returns true when OpenTelemetry tracing is explicitly enabled in
// manifest.backend.observability.tracing. Defaults to **false** (opt-in):
// a missing observability block, missing tracing block, or tracing.enabled:
// false all yield false so projects that don't need distributed tracing
// incur zero OTel build / runtime cost.
func hasOtel(fs *yongol.Fullstack) bool {
	if fs == nil || fs.Manifest == nil {
		return false
	}
	obs := fs.Manifest.Backend.Observability
	if obs == nil || obs.Tracing == nil {
		return false
	}
	return obs.Tracing.Enabled
}
