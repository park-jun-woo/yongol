//ff:func feature=stml-gen type=test control=sequence
//ff:what integer path param이 useMutation에서 Number()로 래핑되는지 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestParamNumberWrapMutation(t *testing.T) {
	page, _ := stmlparser.ParseReader("delete-building.html", strings.NewReader(`<main>
  <button data-action="DeleteBuilding" data-param-id="route.id">삭제</button>
</main>`))
	code := GeneratePage(page, "", GenerateOptions{
		NoBodyOps: map[string]bool{"DeleteBuilding": true},
		PathParamTypes: map[string]map[string]string{
			"DeleteBuilding": {"id": "integer"},
		},
	})
	assertContains(t, code, `id: Number(id)`)
	assertNotContains(t, code, `id: id`)
}
