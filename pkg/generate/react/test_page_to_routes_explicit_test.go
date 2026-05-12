//ff:func feature=gen-react type=test control=sequence
//ff:what pageToRoutes data-route가 detail 추론을 오버라이드하는지 검증

package react

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestPageToRoutes_ExplicitRouteOverridesDetail(t *testing.T) {
	p := stml.PageSpec{
		Name:     "unit-detail",
		FileName: "unit-detail.html",
		Route:    "/buildings/:buildingId/units/:id",
	}
	routes := pageToRoutes(p)
	if len(routes) != 1 {
		t.Fatalf("got %d routes, want 1", len(routes))
	}
	if routes[0].Path != "/buildings/:buildingId/units/:id" {
		t.Errorf("Path = %q, want /buildings/:buildingId/units/:id", routes[0].Path)
	}
	if routes[0].ComponentName != "UnitDetail" {
		t.Errorf("ComponentName = %q, want UnitDetail", routes[0].ComponentName)
	}
}
