//ff:func feature=gen-gogin type=test control=iteration dimension=1 topic=error-envelope
//ff:what TestBlockErrorEnvelope_Defaults — manifest 기본 스냅샷 검증

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
