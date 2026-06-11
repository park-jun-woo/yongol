//ff:func feature=stml-gen type=test control=sequence
//ff:what auth.claims 캡처 방출 — bearer(setAuth 후 setClaim)·cookie(setClaim 만)·클레임 부재 가드 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestGeneratePage_ClaimsCapture(t *testing.T) {
	src := `<main>
  <div data-action="Login" data-capture="access_token -> auth.token, role -> auth.claims.role" data-redirect="/">
    <input data-field="Email" type="email" />
    <button type="submit">로그인</button>
  </div>
</main>`

	t.Run("bearer mode commits token then claim", func(t *testing.T) {
		page, _ := stmlparser.ParseReader("login-page.html", strings.NewReader(src))
		code := GeneratePage(page, "", GenerateOptions{
			APIImportPath: "@/lib/api",
			BearerAuth:    true,
		})
		assertContains(t, code, "import { useAuthStore } from '@/stores/auth'")
		assertContains(t, code, "useAuthStore.getState().setAuth(data.access_token)")
		// Missing claim is skipped, never stored as "undefined".
		assertContains(t, code, "if (data?.role != null) {")
		assertContains(t, code, "useAuthStore.getState().setClaim('role', String(data.role))")
		// The token commit precedes the claim commit.
		if strings.Index(code, "setAuth(data.access_token)") > strings.Index(code, "setClaim('role'") {
			t.Errorf("setAuth should precede setClaim:\n%s", code)
		}
	})

	t.Run("cookie mode commits only the claim", func(t *testing.T) {
		page, _ := stmlparser.ParseReader("login-page.html", strings.NewReader(src))
		code := GeneratePage(page, "", GenerateOptions{
			APIImportPath: "@/lib/api",
			BearerAuth:    false,
		})
		// The token capture stays unemitted (httpOnly cookie carries the
		// session) but the claim is committed from the response body.
		assertNotContains(t, code, "setAuth")
		assertContains(t, code, "import { useAuthStore } from '@/stores/auth'")
		assertContains(t, code, "useAuthStore.getState().setClaim('role', String(data.role))")
		// data-redirect still works alongside the claim commit.
		assertContains(t, code, "navigate('/')")
	})
}
