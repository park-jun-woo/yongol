//ff:func feature=stml-gen type=test control=sequence
//ff:what Login + HasAuthz=false 시 기존 invalidateQueries 동작 유지 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestGenerateLoginPage_NoAuthz_DefaultBehavior(t *testing.T) {
	page, _ := stmlparser.ParseReader("login-page.html", strings.NewReader(`<main>
  <div data-action="Login">
    <input data-field="Email" type="email" />
    <input data-field="Password" type="password" />
    <button type="submit">로그인</button>
  </div>
</main>`))
	code := GeneratePage(page, "", GenerateOptions{
		APIImportPath: "@/lib/api",
		HasAuthz:      false,
	})

	// Default behavior: invalidateQueries
	assertContains(t, code, "queryClient.invalidateQueries()")
	assertContains(t, code, "useQueryClient")

	// Should NOT have token storage or navigate
	assertNotContains(t, code, "localStorage.setItem")
	assertNotContains(t, code, "navigate('/')")
	assertNotContains(t, code, "useNavigate")
}
