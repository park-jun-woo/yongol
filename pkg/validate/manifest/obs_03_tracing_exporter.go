//ff:func feature=validate type=rule control=sequence topic=manifest-observability
//ff:what OBS-003 — backend.observability.tracing.exporter 는 otlp/stdout/noop 중 하나여야 함

package manifest

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// obs03TracingExporter enforces the closed set of supported OTel trace
// exporters so a typo doesn't compile fine and then silently emit no
// spans at runtime. Empty value is accepted (codegen default "noop").
//
// Rule only fires when tracing.enabled is true — operators who leave
// tracing off can still put arbitrary strings in the field without
// blocking validate, since the exporter branch never runs.
func obs03TracingExporter(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.Manifest == nil {
		return nil
	}
	obs := fs.Manifest.Backend.Observability
	if obs == nil || obs.Tracing == nil {
		return nil
	}
	if !obs.Tracing.Enabled {
		return nil
	}
	v := strings.TrimSpace(obs.Tracing.Exporter)
	if v == "" {
		return nil
	}
	switch v {
	case "otlp", "stdout", "noop":
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:    "manifest.yaml",
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: "[OBS-003] backend.observability.tracing.exporter must be one of otlp/stdout/noop (got " + quoted(obs.Tracing.Exporter) + ")",
		Advice:  "exporter 를 otlp(기본 프로덕션), stdout(개발 로컬 확인), noop(비활성) 중 하나로 설정하세요",
	}}
}
