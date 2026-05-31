//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-observability
//ff:what OBS-004 테스트 — 범위 밖 sample_rate 는 ERROR

package manifest

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
