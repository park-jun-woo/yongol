//ff:func feature=stml-gen type=test control=sequence
//ff:what fetch가 소비하는(required ":BuildingID") integer path param은 mutation에서도 가드 없이 Number()로 유지되는지 검증 (BUG-136 무회귀)
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestParamRequiredMutationNoGuard(t *testing.T) {
	// BuildingID is consumed by GetBuilding (required ":BuildingID"), so the
	// DeleteBuilding mutation keeps the unguarded Number(BuildingID) — required
	// params are always present in the matched route.
	page, _ := stmlparser.ParseReader("building-detail.html", strings.NewReader(`<main>
  <article data-fetch="GetBuilding" data-param-building-id="route.BuildingID">
    <h1 data-bind="name"></h1>
  </article>
  <button data-action="DeleteBuilding" data-param-building-id="route.BuildingID" data-redirect="/buildings">삭제</button>
</main>`))
	code := GeneratePage(page, "", GenerateOptions{
		NoBodyOps: map[string]bool{"DeleteBuilding": true},
		PathParamTypes: map[string]map[string]string{
			"GetBuilding":    {"buildingId": "integer"},
			"DeleteBuilding": {"buildingId": "integer"},
		},
	})
	assertContains(t, code, `buildingId: Number(BuildingID) }`)
	assertNotContains(t, code, `BuildingID != null ? Number(BuildingID)`)
	// required-param query also emits no enabled guard
	assertNotContains(t, code, `enabled:`)
}
