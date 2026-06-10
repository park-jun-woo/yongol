//ff:func feature=validate type=test control=sequence topic=manifest-auth
//ff:what SEC-404 테스트 — store=cookie 는 ERROR + backend.auth.mode: cookie 안내

package manifest

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestSec404_CookieStore_Advice(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &pmanifest.ProjectConfig{
			Frontend: pmanifest.Frontend{
				Auth: &pmanifest.FrontendAuth{TokenField: "access_token", Store: "cookie"},
			},
		},
	}
	got := sec404FrontendAuthStoreEnum(fs)
	if len(got) != 1 {
		t.Fatalf("expected 1 diagnostic for store=cookie, got %d: %+v", len(got), got)
	}
	if !strings.Contains(got[0].Message, "[SEC-404]") {
		t.Fatalf("message missing [SEC-404] prefix: %q", got[0].Message)
	}
	if !strings.Contains(got[0].Advice, "backend.auth.mode: cookie") {
		t.Fatalf("advice should point at backend.auth.mode: cookie, got %q", got[0].Advice)
	}
}
