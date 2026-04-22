//ff:func feature=validate type=test control=sequence topic=manifest-observability
//ff:what OBS-004 테스트 — sample_rate 범위 검증

package manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestObs04TracingSampleRate_Valid(t *testing.T) {
	cases := []float64{0.0, 0.05, 0.5, 1.0}
	for _, r := range cases {
		fs := &yongol.Fullstack{
			Manifest: &pmanifest.ProjectConfig{
				Backend: pmanifest.Backend{
					Observability: &pmanifest.Observability{
						Tracing: &pmanifest.ObservabilityTracing{Enabled: true, SampleRate: r},
					},
				},
			},
		}
		if got := obs04TracingSampleRate(fs); len(got) != 0 {
			t.Fatalf("sample_rate=%v expected ok, got %+v", r, got)
		}
	}
}

func TestObs04TracingSampleRate_OutOfRange(t *testing.T) {
	cases := []float64{-0.1, 1.5, 2.0}
	for _, r := range cases {
		fs := &yongol.Fullstack{
			Manifest: &pmanifest.ProjectConfig{
				Backend: pmanifest.Backend{
					Observability: &pmanifest.Observability{
						Tracing: &pmanifest.ObservabilityTracing{Enabled: true, SampleRate: r},
					},
				},
			},
		}
		got := obs04TracingSampleRate(fs)
		if len(got) != 1 {
			t.Fatalf("sample_rate=%v expected 1 diagnostic, got %d", r, len(got))
		}
		if !strings.Contains(got[0].Message, "[OBS-004]") {
			t.Fatalf("message missing [OBS-004] prefix: %q", got[0].Message)
		}
	}
}

func TestObs04TracingSampleRate_Disabled(t *testing.T) {
	// Bad value tolerated while tracing.enabled=false — no runtime impact.
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Observability: &pmanifest.Observability{
					Tracing: &pmanifest.ObservabilityTracing{Enabled: false, SampleRate: 999},
				},
			},
		},
	}
	if got := obs04TracingSampleRate(fs); len(got) != 0 {
		t.Fatalf("expected 0 diagnostics when tracing disabled, got %+v", got)
	}
}
