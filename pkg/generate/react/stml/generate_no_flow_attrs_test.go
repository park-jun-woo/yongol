//ff:func feature=stml-gen type=test control=sequence
//ff:what 흐름 속성 없는 액션은 bearer여도 기존 invalidateQueries 경로 불변 검증
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

	// No data-capture/data-redirect → the default invalidation path stays.
	// "Login" gets no special treatment (hardcode removed in Phase003).
	assertContains(t, code, "queryClient.invalidateQueries()")
	assertContains(t, code, "useQueryClient")

	assertNotContains(t, code, "useAuthStore")
	assertNotContains(t, code, "localStorage.setItem")
	assertNotContains(t, code, "navigate('/')")
	assertNotContains(t, code, "useNavigate")
}
