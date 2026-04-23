//ff:func feature=validate type=util control=sequence topic=manifest-observability
//ff:what resolveMetricsPath — metrics.path 유효 값 결정 (기본 /metrics, 슬래시 누락이면 공백)

package manifest

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// resolveMetricsPath returns the effective scrape path: manifest value when
// explicitly set with a leading "/", else "/metrics". Values failing OBS-001
// (no leading slash) are skipped — OBS-001 already reports them.
func resolveMetricsPath(fs *yongol.Fullstack) string {
	obs := fs.Manifest.Backend.Observability
	if obs == nil || obs.Metrics == nil {
		return "/metrics"
	}
	p := strings.TrimSpace(obs.Metrics.Path)
	if p == "" {
		return "/metrics"
	}
	if !strings.HasPrefix(p, "/") {
		return ""
	}
	return p
}
