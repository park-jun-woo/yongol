//ff:func feature=stml-gen type=test control=sequence
//ff:what 비-Login 액션 + HasAuthz 시 기존 invalidateQueries 유지 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestGenerateNonLoginPage_Authz_NormalInvalidation(t *testing.T) {
	page, _ := stmlparser.ParseReader("create-room-page.html", strings.NewReader(`<main>
  <div data-action="CreateRoom">
    <input data-field="Name" type="text" />
    <button type="submit">생성</button>
  </div>
</main>`))
	code := GeneratePage(page, "", GenerateOptions{
		APIImportPath: "@/lib/api",
		HasAuthz:      true,
	})

	// Normal invalidation behavior
	assertContains(t, code, "queryClient.invalidateQueries()")
	assertContains(t, code, "useQueryClient")

	// Should NOT have token storage or navigate
	assertNotContains(t, code, "localStorage.setItem")
	assertNotContains(t, code, "navigate('/')")
	assertNotContains(t, code, "useNavigate")
}
