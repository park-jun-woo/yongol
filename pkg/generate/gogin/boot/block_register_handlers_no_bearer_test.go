//ff:func feature=gen-gogin type=test control=sequence
//ff:what blockRegisterHandlers — strict-server NewStrictHandler + per-op 미들웨어 + RegisterHandlers
package boot

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBlockRegisterHandlers_NoBearer(t *testing.T) {
	block := blockRegisterHandlers(&yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{}}, "example.com/zenflow")
	body := strings.Join(block.Lines, "\n")
	if strings.Contains(body, "publicOps") || strings.Contains(body, "BearerAuthStrict") {
		t.Errorf("no-bearer project must not emit publicOps / BearerAuthStrict, got:\n%s", body)
	}
	if !strings.Contains(body, "api.NewStrictHandler(srv, []api.StrictMiddlewareFunc{") {
		t.Errorf("must build strict handler, got:\n%s", body)
	}
	if !strings.Contains(body, "api.RegisterHandlers(r, strictHandler)") {
		t.Errorf("must register handlers, got:\n%s", body)
	}
}
