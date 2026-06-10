//ff:func feature=stml-gen type=test control=sequence
//ff:what data-each 행 액션 페이지 생성 — vars 패스스루 mutationFn, 호출부 item.*/route.* 인자, 숫자 item 필드 래핑 생략 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestGenerateRowActionPage(t *testing.T) {
	page, _ := stmlparser.ParseReader("unit-info.html", strings.NewReader(`<main>
  <section data-fetch="GetUnit" data-param-building-id="route.BuildingID">
    <ul data-each="photos">
      <li>
        <span data-bind="caption"></span>
        <button data-action="DeletePhoto"
                data-param-building-id="route.BuildingID"
                data-param-photo-id="item.id">삭제</button>
      </li>
    </ul>
  </section>
</main>`))
	opt := GenerateOptions{
		NoBodyOps: map[string]bool{"DeletePhoto": true},
		PathParamTypes: map[string]map[string]string{
			"GetUnit":     {"buildingId": "integer"},
			"DeletePhoto": {"buildingId": "integer", "photoId": "integer"},
		},
		ResponseArrayItemTypes: map[string]map[string]map[string]string{
			"GetUnit": {"photos": {"id": "integer", "caption": "string"}},
		},
	}
	code := GeneratePage(page, "", opt)

	// The hoisted mutation passes the call-site variables through, typed
	// against the generated api signature (strict tsconfig).
	assertContains(t, code, "mutationFn: (vars: Parameters<typeof api.DeletePhoto>[0]) => api.DeletePhoto(vars)")
	// The row supplies route.* (Number-wrapped) and item.* arguments; the
	// item field is already integer in the response schema → no wrapping.
	assertContains(t, code, "deletePhotoMutation.mutate({ buildingId: Number(BuildingID), photoId: item.id })")
	// The action renders as a trailing row cell inside the table.
	assertContains(t, code, "<TD>")
	assertContains(t, code, "<TH></TH>")
	assertNotContains(t, code, "Number(item.id)")
}
