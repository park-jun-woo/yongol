//ff:func feature=validate type=test control=iteration dimension=1 topic=manifest-observability
//ff:what OBS-004 테스트 — 0.0~1.0 범위의 sample_rate 는 통과

package manifest

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
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
