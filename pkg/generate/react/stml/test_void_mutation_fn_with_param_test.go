//ff:func feature=stml-gen type=test control=sequence
//ff:what NoBodyOps에 포함된 액션의 mutationFn이 () => api.X({ paramArgs }) 형태인지 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestVoidMutationFnWithParam(t *testing.T) {
	page, _ := stmlparser.ParseReader("delete-building.html", strings.NewReader(`<main>
  <button data-action="DeleteBuilding" data-param-id="route.BuildingID">삭제</button>
</main>`))
	code := GeneratePage(page, "", GenerateOptions{
		NoBodyOps: map[string]bool{"DeleteBuilding": true},
	})
	assertContains(t, code, `mutationFn: () => api.DeleteBuilding(`)
	assertNotContains(t, code, `(data) => api.DeleteBuilding`)
	assertNotContains(t, code, `...data`)
}
