//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what tm28RouteSegmentUnused — 라우트 세그먼트가 어떤 data-param-*에도 소비되지 않을 때 WARNING 발화/침묵 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM28RouteSegmentUnused(t *testing.T) {
	cases := []TestTM28RouteSegmentUnusedCase{
		// data-route declares a segment no binding consumes — dead segment.
		{name: "data_route_extra_segment", page: stml.PageSpec{
			FileName: "unit-list.html",
			Route:    "/units/:UnitID/:Extra",
			Fetches: []stml.FetchBlock{{
				OperationID: "ListUnits",
				Params:      []stml.ParamBind{{Name: "unitId", Source: "route.UnitID"}},
			}},
		}, wantCount: 1},
		// Case mismatch counts as unused: useParams() keys are case-sensitive.
		{name: "data_route_case_mismatch_unused", page: stml.PageSpec{
			FileName: "building-edit.html",
			Route:    "/buildings/:buildingId",
			Fetches: []stml.FetchBlock{{
				OperationID: "GetBuilding",
				Params:      []stml.ParamBind{{Name: "buildingId", Source: "route.BuildingID"}},
			}},
		}, wantCount: 1},
		// An unused optional segment (":Name?") is dead too.
		{name: "data_route_unused_optional_segment", page: stml.PageSpec{
			FileName: "unit-info.html",
			Route:    "/units/:UnitID/:PhotoID?",
			Fetches: []stml.FetchBlock{{
				OperationID: "GetUnit",
				Params:      []stml.ParamBind{{Name: "unitId", Source: "route.UnitID"}},
			}},
		}, wantCount: 1},
		// Every segment consumed — silent.
		{name: "data_route_all_segments_consumed", page: stml.PageSpec{
			FileName: "unit-info.html",
			Route:    "/buildings/:BuildingID/units/:UnitID",
			Fetches: []stml.FetchBlock{{
				OperationID: "GetUnit",
				Params: []stml.ParamBind{
					{Name: "buildingId", Source: "route.BuildingID"},
					{Name: "unitId", Source: "route.UnitID"},
				},
			}},
		}, wantCount: 0},
		// Derived route only ever contains consumed params — silent.
		{name: "derived_route_structurally_silent", page: stml.PageSpec{
			FileName: "unit-info.html",
			Fetches: []stml.FetchBlock{{
				OperationID: "GetUnit",
				Params: []stml.ParamBind{
					{Name: "buildingId", Source: "route.BuildingID"},
					{Name: "unitId", Source: "route.UnitID"},
				},
			}},
		}, wantCount: 0},
		// Static explicit route, no params anywhere — silent.
		{name: "static_route_silent", page: stml.PageSpec{
			FileName: "settings.html",
			Route:    "/account/settings",
		}, wantCount: 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runTM28RouteSegmentUnused(t, c)
		})
	}
}
