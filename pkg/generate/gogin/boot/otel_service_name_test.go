//ff:func feature=gen-gogin type=test control=sequence topic=observability
//ff:what otelServiceName — service.name 값 결정 (tracing.service_name → metadata.name → "service")

package boot

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestOtelServiceName(t *testing.T) {
	if got := otelServiceName(nil); got != "service" {
		t.Errorf("nil fs = %q, want service", got)
	}
	if got := otelServiceName(&yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{}}); got != "service" {
		t.Errorf("bare manifest = %q, want service", got)
	}

	metaOnly := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{
		Metadata: pmanifest.Metadata{Name: "zenflow"},
	}}
	if got := otelServiceName(metaOnly); got != "zenflow" {
		t.Errorf("metadata fallback = %q, want zenflow", got)
	}

	explicit := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{
		Metadata: pmanifest.Metadata{Name: "zenflow"},
		Backend: pmanifest.Backend{Observability: &pmanifest.Observability{
			Tracing: &pmanifest.ObservabilityTracing{ServiceName: "billing-svc"},
		}},
	}}
	if got := otelServiceName(explicit); got != "billing-svc" {
		t.Errorf("explicit service_name = %q, want billing-svc", got)
	}
}
