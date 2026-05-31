//ff:func feature=validate type=test control=sequence topic=manifest-observability
//ff:what OBS-003 테스트 — 미허용 exporter 값 (jaeger) 는 ERROR

package manifest

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
