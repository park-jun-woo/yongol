//ff:func feature=gen-ir type=util control=sequence
//ff:what otelEnabled -- manifest.backend.observability.tracing.enabled 여부

package ir

import "github.com/park-jun-woo/yongol/pkg/yongol"

// otelEnabled returns true when manifest.backend.observability.tracing.enabled
// is true.
func otelEnabled(fs *yongol.Fullstack) bool {
	if fs == nil || fs.Manifest == nil {
		return false
	}
	obs := fs.Manifest.Backend.Observability
	if obs == nil || obs.Tracing == nil {
		return false
	}
	return obs.Tracing.Enabled
}
