//ff:func feature=stml-gen type=test control=sequence
//ff:what requiredRouteNames — fetch/중첩fetch 소비 route 세그먼트 집합 수집, 비route 소스 제외 검증

package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRequiredRouteNames(t *testing.T) {
	page := stmlparser.PageSpec{
		Fetches: []stmlparser.FetchBlock{
			{
				OperationID: "GetBuilding",
				Params: []stmlparser.ParamBind{
					{Name: "BuildingID", Source: "route.BuildingID"},
					{Name: "Q", Source: "query.q"}, // non-route, excluded
				},
				NestedFetches: []stmlparser.FetchBlock{
					{OperationID: "ListRooms", Params: []stmlparser.ParamBind{
						{Name: "FloorID", Source: "route.FloorID"},
					}},
				},
			},
		},
	}

	got := requiredRouteNames(page)

	if !got["BuildingID"] {
		t.Errorf("BuildingID should be required")
	}
	if !got["FloorID"] {
		t.Errorf("nested-fetch FloorID should be required")
	}
	if got["Q"] {
		t.Errorf("non-route query.q must not be in required set")
	}
	if len(got) != 2 {
		t.Errorf("required set size = %d, want 2: %v", len(got), got)
	}

	// page without fetches → empty set
	empty := requiredRouteNames(stmlparser.PageSpec{})
	if len(empty) != 0 {
		t.Errorf("page without fetches: want empty set, got %v", empty)
	}
}
