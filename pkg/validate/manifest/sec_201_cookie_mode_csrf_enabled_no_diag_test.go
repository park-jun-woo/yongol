//ff:func feature=validate type=test control=sequence topic=manifest-auth
//ff:what SEC-201 테스트 — cookie 모드 + csrf.enabled=true 는 통과

package manifest

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestSec201_CookieMode_CsrfEnabled_NoDiag(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{
				Auth: &pmanifest.Auth{
					Mode: "cookie",
					Csrf: &pmanifest.CsrfConfig{Enabled: true},
				},
			},
		},
	}
	if got := sec201CookieWithoutCsrf(fs); len(got) != 0 {
		t.Fatalf("cookie + csrf.enabled should not emit SEC-201, got %+v", got)
	}
}
