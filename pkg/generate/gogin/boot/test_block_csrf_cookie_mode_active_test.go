//ff:func feature=gen-gogin type=test control=iteration dimension=1 topic=csrf
//ff:what TestBlockCsrf_CookieMode_Active — cookie 모드에서 csrf 미들웨어 문자열이 모두 포함

package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestBlockCsrf_CookieMode_Active(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Module: "example.com/zenflow",
				Auth: &pmanifest.Auth{
					Mode: "cookie",
					Csrf: &pmanifest.CsrfConfig{
						Enabled:     true,
						CookieName:  "XSRF-TOKEN",
						HeaderName:  "X-XSRF-TOKEN",
						ExemptPaths: []string{"/auth/login", "/auth/refresh"},
					},
				},
			},
		},
	}
	block := blockCsrf(fs, "example.com/zenflow")
	if block.Active == nil || !block.Active(fs) {
		t.Fatalf("cookie mode should report Active()=true")
	}
	body := strings.Join(block.Lines, "\n")
	for _, must := range []string{
		"middleware.Csrf(middleware.CsrfConfig{",
		`"XSRF-TOKEN"`,
		`"X-XSRF-TOKEN"`,
		`"/auth/login"`,
		`"/auth/refresh"`,
		"BACKEND_AUTH_CSRF_ENABLED",
		"HybridBearerSkip: false",
	} {
		if !strings.Contains(body, must) {
			t.Errorf("cookie-mode block missing %q, got:\n%s", must, body)
		}
	}
}
