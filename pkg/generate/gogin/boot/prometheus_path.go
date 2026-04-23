//ff:func feature=gen-gogin type=util control=sequence topic=observability
//ff:what prometheusPath — metrics.path 결정 (미지정 시 "/metrics")

package boot

import "github.com/park-jun-woo/yongol/pkg/yongol"

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
