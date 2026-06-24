//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what blockRegisterHandlersDomained(no bearer) — 도메인별 Group/NewStrictHandler/RegisterHandlers + 공유 srv 검증

package boot

import (
	"strings"
	"testing"
)

func TestBlockRegisterHandlersDomained_NoBearer(t *testing.T) {
	block := blockRegisterHandlers(domainedFS(nil), "example.com/app")
	body := strings.Join(block.Lines, "\n")
	imp := strings.Join(block.Imports, "\n")
	for _, must := range []string{
		`adminGroup := r.Group("/api/admin")`,
		`publicGroup := r.Group("/api")`,
		"api_public.NewStrictHandler(srv, []api_public.StrictMiddlewareFunc{",
		"api_admin.NewStrictHandler(srv, []api_admin.StrictMiddlewareFunc{",
		"api_public.RegisterHandlers(publicGroup, publicStrictHandler)",
		"api_admin.RegisterHandlers(adminGroup, adminStrictHandler)",
		// BUG-142 — per-domain request validator mounted on each route group
		// even without bearer auth.
		"publicValidator, err := middleware.RequestValidatorPublic()",
		"adminValidator, err := middleware.RequestValidatorAdmin()",
		"publicGroup.Use(publicValidator)",
		"adminGroup.Use(adminValidator)",
	} {
		if !strings.Contains(body, must) {
			t.Errorf("missing %q in:\n%s", must, body)
		}
	}
	for _, must := range []string{
		`"example.com/app/internal/api_public"`,
		`"example.com/app/internal/api_admin"`,
		// BUG-142 — validator constructor lives in internal/middleware, so the
		// import is now present regardless of bearer auth.
		`"example.com/app/internal/middleware"`,
	} {
		if !strings.Contains(imp, must) {
			t.Errorf("missing import %q in:\n%s", must, imp)
		}
	}
}
