//ff:func feature=stml-gen type=test control=sequence
//ff:what data-capture+data-redirect 선언 시 store 커밋 + navigate 방출 검증 (bearer)
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestGeneratePage_CaptureRedirect_StoreCommitAndNavigate(t *testing.T) {
	page, _ := stmlparser.ParseReader("login-page.html", strings.NewReader(`<main>
  <div data-action="Login" data-capture="access_token -> auth.token, refresh_token -> auth.refresh" data-redirect="/">
    <input data-field="Email" type="email" />
    <input data-field="Password" type="password" />
    <button type="submit">로그인</button>
  </div>
</main>`))
	code := GeneratePage(page, "", GenerateOptions{
		APIImportPath: "@/lib/api",
		BearerAuth:    true,
	})

	// body only: api function passed directly
	assertContains(t, code, "mutationFn: api.Login")

	// declared captures commit to the session store
	assertContains(t, code, "useAuthStore.getState().setAuth(data.access_token, data.refresh_token)")
	assertContains(t, code, "import { useAuthStore } from '@/stores/auth'")
	assertNotContains(t, code, "localStorage.setItem")

	// declared redirect navigates
	assertContains(t, code, "navigate('/')")
	assertContains(t, code, "useNavigate")
	assertContains(t, code, "from 'react-router-dom'")
	assertContains(t, code, "const navigate = useNavigate()")

	// flow-success path replaces invalidation entirely
	assertNotContains(t, code, "useQueryClient")
	assertNotContains(t, code, "queryClient")
	assertNotContains(t, code, "invalidateQueries")
}
