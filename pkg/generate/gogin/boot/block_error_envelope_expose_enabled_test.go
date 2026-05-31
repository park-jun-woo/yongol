//ff:func feature=gen-gogin type=test control=sequence topic=error-envelope
//ff:what TestBlockErrorEnvelope_ExposeEnabled — expose=true manifest override 확인

package boot

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
