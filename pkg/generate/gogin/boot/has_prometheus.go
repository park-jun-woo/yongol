//ff:func feature=gen-gogin type=util control=sequence topic=observability
//ff:what hasPrometheus — manifest.backend.observability.metrics.enabled (기본 true) 여부

package boot

import "github.com/park-jun-woo/yongol/pkg/yongol"

// hasPrometheus returns true when Prometheus metrics should be wired into
// the generated main.go. Defaults to true (opt-out): a missing manifest
// block or missing enabled flag still yields true so every project ships
// with /metrics out of the box.
func hasPrometheus(fs *yongol.Fullstack) bool {
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

// prometheusPath resolves the scrape endpoint path. Empty / unset manifest
// → "/metrics" default.
func prometheusPath(fs *yongol.Fullstack) string {
	if fs == nil || fs.Manifest == nil {
		return "/metrics"
	}
	obs := fs.Manifest.Backend.Observability
	if obs == nil || obs.Metrics == nil || obs.Metrics.Path == "" {
		return "/metrics"
	}
	return obs.Metrics.Path
}
