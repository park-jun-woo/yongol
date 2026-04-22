//ff:func feature=validate type=rule control=sequence topic=manifest-observability
//ff:what OBS-001 — backend.observability.metrics.path must start with "/"

package manifest

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// obs01MetricsPath enforces that backend.observability.metrics.path, when
// explicitly set, is a valid HTTP route path (leading "/" and non-empty).
// Empty / unset path falls back to "/metrics" at codegen time and needs no
// validation. The rule only fires on explicit malformed values so operators
// can still omit the field safely.
func obs01MetricsPath(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.Manifest == nil {
		return nil
	}
	obs := fs.Manifest.Backend.Observability
	if obs == nil || obs.Metrics == nil {
		return nil
	}
	raw := strings.TrimSpace(obs.Metrics.Path)
	if raw == "" {
		return nil
	}
	if strings.HasPrefix(raw, "/") {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:    "manifest.yaml",
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: "[OBS-001] backend.observability.metrics.path must start with '/' (got " + quoted(obs.Metrics.Path) + ")",
		Advice:  "Set path to an absolute path starting with '/' (e.g. '/metrics' or '/internal/metrics')",
	}}
}
