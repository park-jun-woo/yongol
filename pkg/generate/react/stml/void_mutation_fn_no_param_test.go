//ff:func feature=stml-gen type=test control=sequence
//ff:what NoBodyOps에 포함되고 파라미터가 없는 액션의 mutationFn이 () => api.X() 형태인지 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestVoidMutationFnNoParam(t *testing.T) {
	page, _ := stmlparser.ParseReader("logout-page.html", strings.NewReader(`<main>
  <button data-action="LogoutSession">로그아웃</button>
</main>`))
	code := GeneratePage(page, "", GenerateOptions{
		NoBodyOps: map[string]bool{"LogoutSession": true},
	})
	assertContains(t, code, `mutationFn: () => api.LogoutSession()`)
	assertNotContains(t, code, `(data) => api.LogoutSession`)
}
