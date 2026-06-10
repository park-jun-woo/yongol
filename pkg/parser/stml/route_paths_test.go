//ff:func feature=stml-parse type=test control=sequence
//ff:what TestRoutePaths — 명시 라우트·detail·필수/선택 세그먼트 유도·파일명 유래 패턴 검증

package stml

import (
	"reflect"
	"testing"
)

func TestRoutePaths(t *testing.T) {
	// Explicit data-route wins as-is.
	got := RoutePaths(PageSpec{FileName: "login.html", Route: "/sign-in"})
	if !reflect.DeepEqual(got, []string{"/sign-in"}) {
		t.Errorf("explicit: got %v", got)
	}

	// "-detail" suffix + route.id fetch → pluralized parent + /:id
	// (zenflow's conventional pattern — same result as before derivation).
	got = RoutePaths(PageSpec{
		FileName: "workflow-detail.html",
		Fetches:  []FetchBlock{{OperationID: "GetWorkflow", Params: []ParamBind{{Name: "id", Source: "route.id"}}}},
	})
	if !reflect.DeepEqual(got, []string{"/workflows/:id"}) {
		t.Errorf("detail: got %v", got)
	}

	// "-detail" page that consumes no route.* emits no dead :id segment.
	got = RoutePaths(PageSpec{FileName: "workflow-detail.html"})
	if !reflect.DeepEqual(got, []string{"/workflows"}) {
		t.Errorf("detail without params: got %v", got)
	}

	// Plain page → filename-derived path.
	got = RoutePaths(PageSpec{FileName: "login.html"})
	if !reflect.DeepEqual(got, []string{"/login"}) {
		t.Errorf("plain: got %v", got)
	}

	// Non-detail page with a fetch route param → single derived route
	// (the bare base path is no longer emitted).
	got = RoutePaths(PageSpec{
		FileName: "templates.html",
		Fetches:  []FetchBlock{{OperationID: "GetTemplate", Params: []ParamBind{{Name: "id", Source: "route.id"}}}},
	})
	if !reflect.DeepEqual(got, []string{"/templates/:id"}) {
		t.Errorf("route param: got %v", got)
	}

	// Multiple fetch params → required segments in first-appearance order.
	got = RoutePaths(PageSpec{
		FileName: "unit-info.html",
		Fetches: []FetchBlock{{
			OperationID: "GetUnit",
			Params: []ParamBind{
				{Name: "buildingId", Source: "route.BuildingID"},
				{Name: "unitId", Source: "route.UnitID"},
			},
		}},
	})
	if !reflect.DeepEqual(got, []string{"/unit-info/:BuildingID/:UnitID"}) {
		t.Errorf("multi required: got %v", got)
	}

	// Action-only params become trailing optional segments (":Name?").
	got = RoutePaths(PageSpec{
		FileName: "unit-info.html",
		Fetches: []FetchBlock{{
			OperationID: "GetUnit",
			Params: []ParamBind{
				{Name: "buildingId", Source: "route.BuildingID"},
				{Name: "unitId", Source: "route.UnitID"},
			},
		}},
		Actions: []ActionBlock{{
			OperationID: "DeleteUnitPhoto",
			Params:      []ParamBind{{Name: "photoId", Source: "route.PhotoID"}},
		}},
	})
	if !reflect.DeepEqual(got, []string{"/unit-info/:BuildingID/:UnitID/:PhotoID?"}) {
		t.Errorf("optional: got %v", got)
	}

	// Action-only page → base path + optional segment only.
	got = RoutePaths(PageSpec{
		FileName: "webhooks.html",
		Actions: []ActionBlock{{
			OperationID: "DeleteWebhook",
			Params:      []ParamBind{{Name: "id", Source: "route.id"}},
		}},
	})
	if !reflect.DeepEqual(got, []string{"/webhooks/:id?"}) {
		t.Errorf("action only: got %v", got)
	}
}
