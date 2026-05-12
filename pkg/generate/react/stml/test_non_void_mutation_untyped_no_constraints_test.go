//ff:func feature=stml-gen type=test control=sequence
//ff:what body+path 분기에서 constraints 없을 때 data 타입이 명시되지 않는지 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestNonVoidMutationUntypedDataWithoutConstraints(t *testing.T) {
	page, _ := stmlparser.ParseReader("update-building.html", strings.NewReader(`<main>
  <form data-action="UpdateBuilding" data-param-id="route.BuildingID">
    <input name="name" />
    <button type="submit">수정</button>
  </form>
</main>`))
	code := GeneratePage(page, "", GenerateOptions{
		PathParamTypes: map[string]map[string]string{
			"UpdateBuilding": {"id": "integer"},
		},
	})
	assertContains(t, code, `mutationFn: (data) => api.UpdateBuilding({ ...data,`)
	assertNotContains(t, code, `z.infer`)
}
