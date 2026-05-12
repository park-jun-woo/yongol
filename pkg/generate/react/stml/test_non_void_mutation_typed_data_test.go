//ff:func feature=stml-gen type=test control=sequence
//ff:what body+path 분기에서 constraints 존재 시 data에 z.infer<typeof schema> 타입이 명시되는지 검증
package stml

import (
	"strings"
	"testing"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestNonVoidMutationTypedData(t *testing.T) {
	page, _ := stmlparser.ParseReader("update-building.html", strings.NewReader(`<main>
  <form data-action="UpdateBuilding" data-param-id="route.BuildingID">
    <input name="name" />
    <button type="submit">수정</button>
  </form>
</main>`))
	code := GeneratePage(page, "", GenerateOptions{
		RequestConstraints: map[string]map[string]oapiparser.FieldConstraint{
			"UpdateBuilding": {
				"name": {Type: "string", Required: true},
			},
		},
		PathParamTypes: map[string]map[string]string{
			"UpdateBuilding": {"id": "integer"},
		},
	})
	assertContains(t, code, `(data: z.infer<typeof updateBuildingSchema>) => api.UpdateBuilding({ ...data,`)
	assertNotContains(t, code, `mutationFn: (data) => api.UpdateBuilding`)
}
