//ff:func feature=stml-gen type=test control=sequence
//ff:what cookie 모드(BearerAuth=false)에서 캡처 미방출·redirect만 동작 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestGeneratePage_CookieMode_NoCaptureRedirectOnly(t *testing.T) {
	page, _ := stmlparser.ParseReader("login-page.html", strings.NewReader(`<main>
  <div data-action="Login" data-capture="access_token -> auth.token" data-redirect="/">
    <input data-field="Email" type="email" />
    <button type="submit">로그인</button>
  </div>
</main>`))
	code := GeneratePage(page, "", GenerateOptions{
		APIImportPath: "@/lib/api",
		BearerAuth:    false,
	})

	// cookie mode: no store, no capture commit (httpOnly cookie carries the
	// session) — TM-24 diagnoses the stale capture at validate time.
	assertNotContains(t, code, "useAuthStore")
	assertNotContains(t, code, "setAuth")

	// data-redirect still works in cookie mode.
	assertContains(t, code, "navigate('/')")
	assertContains(t, code, "useNavigate")
	assertNotContains(t, code, "invalidateQueries")
}
