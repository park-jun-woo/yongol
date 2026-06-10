//ff:func feature=stml-gen type=test control=sequence
//ff:what 페이지명 data-redirect + data-redirect-params 선언 시 부재 가드 + 동적 navigate 방출 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestGeneratePage_DynamicRedirect_NavigateWithResponseField(t *testing.T) {
	page, _ := stmlparser.ParseReader("contract-new.html", strings.NewReader(`<main>
  <div data-action="CreateContract" data-redirect="contract-edit" data-redirect-params="id -> ContractID">
    <input data-field="Title" type="text" />
    <button type="submit">등록</button>
  </div>
</main>`))
	code := GeneratePage(page, "", GenerateOptions{
		APIImportPath: "@/lib/api",
		RoutePatterns: map[string]string{"contract-edit": "/contract-edit/:ContractID"},
	})

	// onSuccess receives the 2xx response to substitute from
	assertContains(t, code, "onSuccess: (data) => {")

	// defensive navigate (Phase003-style guard): a 2xx response missing the
	// substituted field aborts the navigate and surfaces through the
	// always-present error state (page-flow Phase004)
	assertContains(t, code, "if (data?.id == null) {")
	assertContains(t, code, "setCreateContractError('Unexpected response: missing id')")
	assertContains(t, code, "return")

	// the page-name reference resolves to the target route with the
	// response field substituted
	assertContains(t, code, "navigate(`/contract-edit/${data.id}`)")
	assertContains(t, code, "useNavigate")
	assertContains(t, code, "const navigate = useNavigate()")

	// flow-success path replaces invalidation entirely
	assertNotContains(t, code, "useQueryClient")
	assertNotContains(t, code, "invalidateQueries")
}
