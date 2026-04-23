//ff:func feature=validate type=rule control=iteration dimension=1 topic=manifest-observability
//ff:what OBS-002 — metrics.path must not collide with an OpenAPI-defined path

package manifest

import (
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
				Advice:  "Change metrics.path to a path not present in OpenAPI (e.g. '/internal/metrics'), or remove that path from the OpenAPI document",
			}}
		}
	}
	return nil
}
