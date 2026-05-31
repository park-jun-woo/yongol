//ff:func feature=gen-gogin type=test control=sequence topic=error-envelope
//ff:what blockErrorEnvelope — middleware.ErrorEnvelopeMiddleware 등록 (Phase004)
package boot

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBlockErrorEnvelope_Default(t *testing.T) {
	block := blockErrorEnvelope(&yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{}}, "example.com/zenflow")
	body := strings.Join(block.Lines, "\n")
	if !strings.Contains(body, `middleware.ExposeInternalError = envBool("BACKEND_ERROR_EXPOSE_INTERNAL_ERROR", false)`) {
		t.Errorf("default expose should be false, got:\n%s", body)
	}
	if !strings.Contains(body, "r.Use(middleware.ErrorEnvelopeMiddleware())") {
		t.Errorf("must register envelope middleware, got:\n%s", body)
	}
}
