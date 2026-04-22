//ff:func feature=validate type=rule control=sequence topic=manifest-observability
//ff:what OBS-001 — backend.observability.metrics.path 가 "/" 로 시작해야 함

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
		Advice:  "path 를 '/metrics' 또는 '/internal/metrics' 처럼 '/' 로 시작하는 절대 경로로 지정하세요",
	}}
}
