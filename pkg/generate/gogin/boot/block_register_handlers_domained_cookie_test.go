//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what blockRegisterHandlersDomained(mixed) — cookie 도메인 → CookieAuthStrict<Title>, bearer 도메인 → BearerAuthStrict<Title> 검증

package boot

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBlockRegisterHandlersDomained_CookieMode(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Backend: manifest.Backend{
				Module:     "example.com/app",
				Middleware: []string{"bearerAuth"},
				Auth:       &manifest.Auth{Mode: "cookie"},
			},
			Domains: map[string]manifest.DomainConfig{
				"public": {RoutePrefix: "/api"},                           // inherits cookie
				"admin":  {RoutePrefix: "/api/admin", AuthMode: "bearer"}, // override bearer
			},
		},
		DomainOpenAPIDocs: map[string]*openapi3.T{
			"public": {Paths: openapi3.NewPaths()},
			"admin":  {Paths: openapi3.NewPaths()},
		},
	}
	block := blockRegisterHandlers(fs, "example.com/app")
	body := strings.Join(block.Lines, "\n")
	for _, must := range []string{
		"middleware.CookieAuthStrictPublic(publicPublicOps)",
		"middleware.BearerAuthStrictAdmin(adminPublicOps)",
	} {
		if !strings.Contains(body, must) {
			t.Errorf("missing %q in:\n%s", must, body)
		}
	}
}
