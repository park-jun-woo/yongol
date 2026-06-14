//ff:func feature=stml-gen type=test control=sequence
//ff:what action-only(optional ":id?") integer path param이 useMutation에서 null 가드된 Number()로 방출되는지 검증 (BUG-136)
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
	// route.id is action-only → optional segment ":id?", so an absent segment
	// must not send Number(undefined)===NaN (BUG-136).
	assertContains(t, code, `id: id != null ? Number(id) : undefined`)
	assertNotContains(t, code, `id: Number(id) }`)
}
