//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what tm27RouteParamMissing — 소비 route.* 가 해석 라우트 세그먼트에 없을 때 ERROR 발화/침묵 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM27RouteParamMissing(t *testing.T) {
	cases := []TestTM27RouteParamMissingCase{
		// Explicit data-route whose segment name does not match the
		// consumed param — the BUG-112 shape (":id" slot vs route.UnitID).
		{name: "data_route_segment_name_mismatch", page: stml.PageSpec{
			FileName: "unit-info.html",
			Route:    "/units/:id",
			Fetches: []stml.FetchBlock{{
				OperationID: "GetUnit",
				Params:      []stml.ParamBind{{Name: "unitId", Source: "route.UnitID"}},
			}},
		}, wantCount: 1},
		// Case mismatch is a mismatch: useParams() keys are case-sensitive.
		{name: "data_route_case_mismatch", page: stml.PageSpec{
			FileName: "building-edit.html",
			Route:    "/buildings/:buildingId",
			Fetches: []stml.FetchBlock{{
				OperationID: "GetBuilding",
				Params:      []stml.ParamBind{{Name: "buildingId", Source: "route.BuildingID"}},
			}},
		}, wantCount: 1},
		// Arity mismatch: one slot, three consumed params → 2 missing.
		{name: "data_route_arity_mismatch", page: stml.PageSpec{
			FileName: "unit-info.html",
			Route:    "/unit-info/:BuildingID",
			Fetches: []stml.FetchBlock{{
				OperationID: "GetUnit",
				Params: []stml.ParamBind{
					{Name: "buildingId", Source: "route.BuildingID"},
					{Name: "unitId", Source: "route.UnitID"},
				},
			}},
			Actions: []stml.ActionBlock{{
				OperationID: "DeletePhoto",
				Params:      []stml.ParamBind{{Name: "photoId", Source: "route.PhotoID"}},
			}},
		}, wantCount: 2},
		// Explicit data-route with matching segments — silent, optional
		// marker ("?") stripped before comparison.
		{name: "data_route_all_segments_match", page: stml.PageSpec{
			FileName: "unit-info.html",
			Route:    "/buildings/:BuildingID/units/:UnitID/:PhotoID?",
			Fetches: []stml.FetchBlock{{
				OperationID: "GetUnit",
				Params: []stml.ParamBind{
					{Name: "buildingId", Source: "route.BuildingID"},
					{Name: "unitId", Source: "route.UnitID"},
				},
			}},
			Actions: []stml.ActionBlock{{
				OperationID: "DeletePhoto",
				Params:      []stml.ParamBind{{Name: "photoId", Source: "route.PhotoID"}},
			}},
		}, wantCount: 0},
		// Derived route (no data-route) is built from the consumed set —
		// structurally satisfied; the rule stays a regression guard.
		{name: "derived_route_structurally_silent", page: stml.PageSpec{
			FileName: "unit-info.html",
			Fetches: []stml.FetchBlock{{
				OperationID: "GetUnit",
				Params: []stml.ParamBind{
					{Name: "buildingId", Source: "route.BuildingID"},
					{Name: "unitId", Source: "route.UnitID"},
				},
			}},
			Actions: []stml.ActionBlock{{
				OperationID: "DeletePhoto",
				Params:      []stml.ParamBind{{Name: "photoId", Source: "route.PhotoID"}},
			}},
		}, wantCount: 0},
		// No route.* consumption — nothing to check.
		{name: "no_route_consumption_silent", page: stml.PageSpec{
			FileName: "login.html",
			Route:    "/login",
		}, wantCount: 0},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runTM27RouteParamMissing(t, c)
		})
	}
}
