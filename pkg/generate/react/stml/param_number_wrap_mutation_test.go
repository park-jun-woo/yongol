//ff:func feature=stml-gen type=test control=sequence
//ff:what action-only(optional ":id?") integer path param의 arg는 평문 Number()(BUG-137)이고 트리거 버튼이 부재 시 disabled되는지(BUG-136) 검증
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
	// route.id is action-only → optional segment ":id?". The arg stays plain
	// Number(id) for type fidelity (BUG-137); the absent-segment NaN is blocked
	// by the trigger guard (BUG-136), not by mutilating the arg.
	assertContains(t, code, `id: Number(id) }`)
	assertNotContains(t, code, `id != null ? Number(id) : undefined`)
	// The button trigger is disabled while the optional param is absent.
	assertContains(t, code, `id == null`)
}
