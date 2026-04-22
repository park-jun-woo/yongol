//ff:func feature=gen-gogin type=test control=sequence topic=observability
//ff:what blockRouter — tracing 활성 시 otelgin.Middleware 등록 회귀 방지

package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestBlockRouter_Default(t *testing.T) {
	// No tracing → no otelgin import, no middleware registration.
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{Backend: pmanifest.Backend{}},
	}
	block := blockRouter(fs, "example.com/zenflow")
	body := strings.Join(block.Lines, "\n")
	imports := strings.Join(block.Imports, "\n")
	if !strings.Contains(body, "gin.Default()") {
		t.Fatalf("router block must create gin engine, got:\n%s", body)
	}
	if strings.Contains(body, "otelgin") {
		t.Fatalf("non-tracing router must NOT register otelgin, got:\n%s", body)
	}
	if strings.Contains(imports, "otelgin") {
		t.Fatalf("non-tracing router must NOT import otelgin, got:\n%s", imports)
	}
}

func TestBlockRouter_TracingEnabled(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Observability: &pmanifest.Observability{
					Tracing: &pmanifest.ObservabilityTracing{Enabled: true},
				},
			},
		},
	}
	block := blockRouter(fs, "example.com/zenflow")
	body := strings.Join(block.Lines, "\n")
	imports := strings.Join(block.Imports, "\n")
	if !strings.Contains(body, "otelgin.Middleware(otelServiceName)") {
		t.Fatalf("tracing-enabled router must register otelgin.Middleware, got:\n%s", body)
	}
	if !strings.Contains(imports, "otelgin") {
		t.Fatalf("missing otelgin import, got:\n%s", imports)
	}
}
