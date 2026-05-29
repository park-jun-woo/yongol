//ff:func feature=stml-gen type=test control=sequence
//ff:what body only 액션에서 mutationFn이 api 함수 직접 참조 형태이고 arrow function이 아닌지 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestBodyOnlyMutationDirectRef(t *testing.T) {
	page, _ := stmlparser.ParseReader("create-item.html", strings.NewReader(`<main>
  <form data-action="CreateItem">
    <input name="title" />
    <button type="submit">생성</button>
  </form>
</main>`))
	code := GeneratePage(page, "")
	assertContains(t, code, `mutationFn: api.CreateItem`)
	assertNotContains(t, code, `(data) => api.CreateItem(data)`)
	assertNotContains(t, code, `() => api.CreateItem`)
}
