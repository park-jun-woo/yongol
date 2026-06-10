//ff:func feature=gen-react type=test control=sequence
//ff:what pageToRoutes stml.RoutePaths 유도 경로(다중 필수 + optional 세그먼트) 소비 검증

package react

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestPageToRoutes_DerivedFromRouteConsumption(t *testing.T) {
	p := stml.PageSpec{
		Name:     "unit-info",
		FileName: "unit-info.html",
		Fetches: []stml.FetchBlock{{
			OperationID: "GetUnit",
			Params: []stml.ParamBind{
				{Name: "buildingId", Source: "route.BuildingID"},
				{Name: "unitId", Source: "route.UnitID"},
			},
		}},
		Actions: []stml.ActionBlock{{
			OperationID: "DeleteUnitPhoto",
			Params:      []stml.ParamBind{{Name: "photoId", Source: "route.PhotoID"}},
		}},
	}
	routes := pageToRoutes(p)
	if len(routes) != 1 {
		t.Fatalf("got %d routes, want 1", len(routes))
	}
	if routes[0].Path != "/unit-info/:BuildingID/:UnitID/:PhotoID?" {
		t.Errorf("Path = %q, want /unit-info/:BuildingID/:UnitID/:PhotoID?", routes[0].Path)
	}
	if routes[0].ComponentName != "UnitInfo" {
		t.Errorf("ComponentName = %q, want UnitInfo", routes[0].ComponentName)
	}
	if routes[0].ImportPath != "./pages/unit-info" {
		t.Errorf("ImportPath = %q, want ./pages/unit-info", routes[0].ImportPath)
	}
}
