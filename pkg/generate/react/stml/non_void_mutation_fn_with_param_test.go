//ff:func feature=stml-gen type=test control=sequence
//ff:what NoBodyOps에 미포함된 액션의 mutationFn이 (data) => api.X({ ...data, paramArgs }) 형태인지 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestNonVoidMutationFnWithParam(t *testing.T) {
	page, _ := stmlparser.ParseReader("update-building.html", strings.NewReader(`<main>
  <form data-action="UpdateBuilding" data-param-id="route.BuildingID">
    <input name="name" />
    <button type="submit">수정</button>
  </form>
</main>`))
	code := GeneratePage(page, "")
	assertContains(t, code, `mutationFn: (data) => api.UpdateBuilding({ ...data,`)
	assertNotContains(t, code, `mutationFn: () =>`)
}
