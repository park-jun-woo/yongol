//ff:func feature=stml-gen type=test control=sequence
//ff:what Login 외 op도 캡처 선언만으로 동일한 store 커밋 방출 검증 (비대칭 해소)
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestGeneratePage_CaptureNonLoginOp_SameEmission(t *testing.T) {
	page, _ := stmlparser.ParseReader("twofa-page.html", strings.NewReader(`<main>
  <div data-action="Verify2FA" data-capture="access_token -> auth.token" data-redirect="/dashboard">
    <input data-field="Code" type="text" />
    <button type="submit">확인</button>
  </div>
</main>`))
	code := GeneratePage(page, "", GenerateOptions{
		APIImportPath: "@/lib/api",
		BearerAuth:    true,
	})

	// The op name is irrelevant — the data-capture declaration decides.
	assertContains(t, code, "useAuthStore.getState().setAuth(data.access_token)")
	assertContains(t, code, "navigate('/dashboard')")
	assertNotContains(t, code, "invalidateQueries")
}
