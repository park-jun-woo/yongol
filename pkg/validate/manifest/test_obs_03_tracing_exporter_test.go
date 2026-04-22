//ff:func feature=validate type=test control=sequence topic=manifest-observability
//ff:what OBS-003 테스트 — 허용된 exporter 값은 조용히, 그 외는 오류

package manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestObs03TracingExporter_Allowed(t *testing.T) {
	cases := []string{"otlp", "stdout", "noop", ""}
	for _, v := range cases {
		t.Run("exporter="+v, func(t *testing.T) {
			fs := &yongol.Fullstack{
				Manifest: &pmanifest.ProjectConfig{
					Backend: pmanifest.Backend{
						Observability: &pmanifest.Observability{
							Tracing: &pmanifest.ObservabilityTracing{Enabled: true, Exporter: v},
						},
					},
				},
			}
			if got := obs03TracingExporter(fs); len(got) != 0 {
				t.Fatalf("expected 0 diagnostics, got %d: %+v", len(got), got)
			}
		})
	}
}

func TestObs03TracingExporter_Disabled(t *testing.T) {
	// Bogus value is tolerated while tracing.enabled is false — the
	// generator never reads exporter in that case.
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Observability: &pmanifest.Observability{
					Tracing: &pmanifest.ObservabilityTracing{Enabled: false, Exporter: "bogus"},
				},
			},
		},
	}
	if got := obs03TracingExporter(fs); len(got) != 0 {
		t.Fatalf("expected 0 diagnostics when tracing disabled, got: %+v", got)
	}
}

func TestObs03TracingExporter_Negative(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Observability: &pmanifest.Observability{
					Tracing: &pmanifest.ObservabilityTracing{Enabled: true, Exporter: "jaeger"},
				},
			},
		},
	}
	got := obs03TracingExporter(fs)
	if len(got) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d", len(got))
	}
	if !strings.Contains(got[0].Message, "[OBS-003]") {
		t.Fatalf("message missing [OBS-003] prefix: %q", got[0].Message)
	}
}
