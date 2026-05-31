//ff:func feature=validate type=test control=sequence topic=manifest-auth
//ff:what SEC-201 테스트 — cookie 모드 + csrf.enabled=false negative (ERROR)

package manifest

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestSec201_CookieMode_CsrfDisabled_Error(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Auth: &pmanifest.Auth{
					Mode: "cookie",
					Csrf: &pmanifest.CsrfConfig{Enabled: false},
				},
			},
		},
	}
	got := sec201CookieWithoutCsrf(fs)
	if len(got) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d: %+v", len(got), got)
	}
	if !strings.Contains(got[0].Message, "[SEC-201]") {
		t.Fatalf("message missing [SEC-201] prefix: %q", got[0].Message)
	}
}
