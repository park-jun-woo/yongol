//ff:func feature=gen-gogin type=test control=sequence topic=error-envelope
//ff:what blockErrorEnvelope 기본/manifest override 스냅샷

package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestBlockErrorEnvelope_Defaults(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{Module: "example.com/zenflow"},
		},
	}
	block := blockErrorEnvelope(fs, "example.com/zenflow")
	if block.Name != "error-envelope" {
		t.Fatalf("unexpected name %q", block.Name)
	}
	body := strings.Join(block.Lines, "\n")
	for _, must := range []string{
		`middleware.ExposeInternalError = envBool("BACKEND_ERROR_EXPOSE_INTERNAL_ERROR", false)`,
		`r.Use(middleware.ErrorEnvelopeMiddleware())`,
	} {
		if !strings.Contains(body, must) {
			t.Errorf("error-envelope block missing %q; body:\n%s", must, body)
		}
	}
}

func TestBlockErrorEnvelope_ExposeEnabled(t *testing.T) {
	expose := true
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Module: "example.com/zenflow",
				Error: &pmanifest.ErrorConfig{
					ExposeInternalError: &expose,
				},
			},
		},
	}
	block := blockErrorEnvelope(fs, "example.com/zenflow")
	body := strings.Join(block.Lines, "\n")
	if !strings.Contains(body, `envBool("BACKEND_ERROR_EXPOSE_INTERNAL_ERROR", true)`) {
		t.Errorf("expected expose default=true, got:\n%s", body)
	}
}
