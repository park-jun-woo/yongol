//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what blockRegisterHandlers — strict-server NewStrictHandler + per-op 미들웨어 + RegisterHandlers

package boot

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
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

func TestBlockRegisterHandlers_WithBearer(t *testing.T) {
	doc := buildDoc([]opSpec{
		{path: "/login", method: "POST", opID: "Login", sec: &openapi3.SecurityRequirements{}},
		{path: "/me", method: "GET", opID: "Me"},
	}, true)
	fs := &yongol.Fullstack{
		OpenAPIDoc: doc,
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{Middleware: []string{"bearerAuth"}},
		},
	}
	block := blockRegisterHandlers(fs, "example.com/zenflow")
	body := strings.Join(block.Lines, "\n")
	if !strings.Contains(body, "publicOps := map[string]bool{") {
		t.Errorf("bearer project must emit publicOps map, got:\n%s", body)
	}
	if !strings.Contains(body, `"Login": true,`) {
		t.Errorf("Login (security: []) should be a public op, got:\n%s", body)
	}
	if !strings.Contains(body, "middleware.BearerAuthStrict(publicOps)") {
		t.Errorf("must inject BearerAuthStrict, got:\n%s", body)
	}
	if !strings.Contains(strings.Join(block.Imports, "\n"), `"example.com/zenflow/internal/middleware"`) {
		t.Errorf("must import middleware, got:\n%v", block.Imports)
	}
}
