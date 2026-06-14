//ff:func feature=stml-gen type=test control=sequence
//ff:what 흐름 속성·fetch 없는 액션은 keyless invalidateQueries()를 방출하지 않음 (BUG-132 132-3)
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestGeneratePage_NoFlowAttrs_DefaultInvalidate(t *testing.T) {
	page, _ := stmlparser.ParseReader("login-page.html", strings.NewReader(`<main>
  <div data-action="Login">
    <input data-field="Email" type="email" />
    <input data-field="Password" type="password" />
    <button type="submit">로그인</button>
  </div>
</main>`))
	code := GeneratePage(page, "", GenerateOptions{
		APIImportPath: "@/lib/api",
		BearerAuth:    true,
	})

	// No data-capture/data-redirect and no page-level fetch → there is no
	// query key to invalidate, so the keyless invalidateQueries() that would
	// wipe the whole cache is gone (BUG-132 132-3) and queryClient is not
	// imported. "Login" gets no special treatment (hardcode removed Phase003).
	assertNotContains(t, code, "invalidateQueries")
	assertNotContains(t, code, "useQueryClient")

	assertNotContains(t, code, "useAuthStore")
	assertNotContains(t, code, "localStorage.setItem")
	assertNotContains(t, code, "navigate('/')")
	assertNotContains(t, code, "useNavigate")
}
