//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-observability
//ff:what obs03TracingExporter — tracing.exporter가 otlp/stdout/noop 중 하나인지 검증

package manifest

import (
	"testing"

	pm "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestObs03TracingExporter(t *testing.T) {
	mk := func(enabled bool, exporter string) *yongol.Fullstack {
		return &yongol.Fullstack{Manifest: &pm.ProjectConfig{Backend: pm.Backend{
			Observability: &pm.Observability{Tracing: &pm.ObservabilityTracing{Enabled: enabled, Exporter: exporter}},
		}}}
	}
	cases := []TestObs03TracingExporterCase{
		{name: "nil_fs", fs: nil, wantCount: 0},
		{name: "nil_manifest", fs: &yongol.Fullstack{}, wantCount: 0},
		{name: "nil_obs", fs: &yongol.Fullstack{Manifest: &pm.ProjectConfig{}}, wantCount: 0},
		{name: "disabled", fs: mk(false, "invalid"), wantCount: 0},
		{name: "empty_exporter", fs: mk(true, ""), wantCount: 0},
		{name: "otlp", fs: mk(true, "otlp"), wantCount: 0},
		{name: "stdout", fs: mk(true, "stdout"), wantCount: 0},
		{name: "noop", fs: mk(true, "noop"), wantCount: 0},
		{name: "invalid", fs: mk(true, "jaeger"), wantCount: 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runObs03TracingExporter(t, c)
		})
	}
}
