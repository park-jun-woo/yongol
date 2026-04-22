//ff:func feature=validate type=rule control=sequence topic=manifest-observability
//ff:what OBS-002 — metrics.path 가 OpenAPI 정의 path 와 충돌하면 안 됨

package manifest

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// obs02MetricsPathNotOpenAPI prevents the Prometheus scrape endpoint from
// shadowing a real API route. Comparison is literal on the configured path
// (default "/metrics" when unset) against every OpenAPI path key.
//
// OpenAPI paths containing templates (e.g. "/users/{id}") cannot collide
// with a static metrics path, so the rule only fires on exact string match.
// That keeps the rule intent-preserving — it catches the one class of bug
// operators actually hit: adding "/metrics" to the OpenAPI doc by mistake.
func obs02MetricsPathNotOpenAPI(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.Manifest == nil {
		return nil
	}
	path := resolveMetricsPath(fs)
	if path == "" {
		return nil
	}
	if fs.OpenAPIDoc == nil || fs.OpenAPIDoc.Paths == nil {
		return nil
	}
	for apiPath := range fs.OpenAPIDoc.Paths.Map() {
		if apiPath == path {
			return []diagnostic.Diagnostic{{
				File:    "manifest.yaml",
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: "[OBS-002] backend.observability.metrics.path " + quoted(path) + " collides with an OpenAPI path of the same name",
				Advice:  "metrics.path 를 OpenAPI 에 없는 경로 (예: '/internal/metrics') 로 변경하거나 OpenAPI 에서 해당 path 를 제거하세요",
			}}
		}
	}
	return nil
}

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
