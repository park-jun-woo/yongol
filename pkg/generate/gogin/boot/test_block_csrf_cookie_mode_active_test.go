//ff:func feature=gen-gogin type=test control=iteration dimension=1 topic=csrf
//ff:what TestBlockCsrf_CookieMode_Active — cookie 모드에서 csrf 미들웨어 문자열이 모두 포함

package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestBlockCsrf_CookieMode_Active(t *testing.T) {
	raw := &pmanifest.Auth{
		Mode: "cookie",
		Csrf: &pmanifest.CsrfConfig{
			Enabled:     true,
			CookieName:  "XSRF-TOKEN",
			HeaderName:  "X-XSRF-TOKEN",
			ExemptPaths: []string{"/auth/login", "/auth/refresh"},
		},
	}
	a := prepared.Auth{Present: true, Mode: "cookie", Raw: raw}
	block := blockCsrf(a, "example.com/zenflow")
	if block.Active != nil {
		t.Fatalf("cookie mode with csrf enabled should leave Active nil (always active)")
	}
	if len(block.Lines) == 0 {
		t.Fatalf("cookie mode with csrf enabled should emit lines")
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
