//ff:func feature=gen-ir type=util control=sequence
//ff:what prometheusEnabled -- manifest.backend.observability.metrics.enabled 여부 (기본 true)

package ir

import "github.com/park-jun-woo/yongol/pkg/yongol"

// prometheusEnabled returns true when Prometheus metrics are active.
// Defaults to true (opt-out).
func prometheusEnabled(fs *yongol.Fullstack) bool {
	if fs == nil || fs.Manifest == nil {
		return true
	}
	obs := fs.Manifest.Backend.Observability
	if obs == nil || obs.Metrics == nil {
		return true
	}
	if obs.Metrics.Enabled == nil {
		return true
	}
	return *obs.Metrics.Enabled
}
