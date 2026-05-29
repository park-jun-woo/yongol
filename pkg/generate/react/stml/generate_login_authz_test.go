//ff:func feature=stml-gen type=test control=sequence
//ff:what Login + HasAuthz 시 토큰 저장 + navigate('/') + body only 직접 참조 방출 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestGenerateLoginPage_Authz_TokenStore(t *testing.T) {
	page, _ := stmlparser.ParseReader("login-page.html", strings.NewReader(`<main>
  <div data-action="Login">
    <input data-field="Email" type="email" />
    <input data-field="Password" type="password" />
    <button type="submit">로그인</button>
  </div>
</main>`))
	code := GeneratePage(page, "", GenerateOptions{
		APIImportPath: "@/lib/api",
		HasAuthz:      true,
	})

	// body only: api function passed directly
	assertContains(t, code, "mutationFn: api.Login")
	assertNotContains(t, code, "(data) => api.Login(data)")

	// Token storage in onSuccess
	assertContains(t, code, "localStorage.setItem('access_token', data.access_token)")
	assertContains(t, code, "localStorage.setItem('refresh_token', data.refresh_token)")
	assertContains(t, code, "navigate('/')")

	// useNavigate import
	assertContains(t, code, "useNavigate")
	assertContains(t, code, "from 'react-router-dom'")
	assertContains(t, code, "const navigate = useNavigate()")

	// Should NOT have queryClient (only Login action, no invalidation needed)
	assertNotContains(t, code, "useQueryClient")
	assertNotContains(t, code, "queryClient")
	assertNotContains(t, code, "invalidateQueries")
}
