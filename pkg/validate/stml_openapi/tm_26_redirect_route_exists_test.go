//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM26RedirectRouteExists — 미해석 경로 ERROR·인덱스 허용·파일명/패턴 매칭 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM26RedirectRouteExists(t *testing.T) {
	pages := []stml.PageSpec{
		{FileName: "login.html"},
		{FileName: "dashboard.html"},
		{
			FileName: "workflow-detail.html",
			Fetches: []stml.FetchBlock{{
				OperationID: "GetWorkflow",
				Params:      []stml.ParamBind{{Name: "id", Source: "route.id"}},
			}},
		},
		{FileName: "settings.html", Route: "/account/settings"},
	}

	// "/" is always allowed (index route, emitted by Phase005).
	if d := tm26RedirectRouteExists(stml.ActionBlock{OperationID: "Login", Redirect: "/"}, "login.html", pages); len(d) != 0 {
		t.Errorf("index: expected 0 diagnostics, got %+v", d)
	}

	// Filename-derived route resolves.
	if d := tm26RedirectRouteExists(stml.ActionBlock{OperationID: "Login", Redirect: "/dashboard"}, "login.html", pages); len(d) != 0 {
		t.Errorf("dashboard: expected 0 diagnostics, got %+v", d)
	}

	// Detail pattern with a concrete :id segment resolves.
	if d := tm26RedirectRouteExists(stml.ActionBlock{OperationID: "CreateWorkflow", Redirect: "/workflows/3"}, "p.html", pages); len(d) != 0 {
		t.Errorf("detail: expected 0 diagnostics, got %+v", d)
	}

	// Explicit data-route resolves.
	if d := tm26RedirectRouteExists(stml.ActionBlock{OperationID: "Login", Redirect: "/account/settings"}, "login.html", pages); len(d) != 0 {
		t.Errorf("explicit route: expected 0 diagnostics, got %+v", d)
	}

	// Unresolved path → 1 ERROR.
	got := tm26RedirectRouteExists(stml.ActionBlock{OperationID: "Login", Redirect: "/nope"}, "login.html", pages)
	if countDiag(got, "[TM-26]") != 1 {
		t.Errorf("unresolved: expected 1 TM-26, got %+v", got)
	}

	// No redirect → no diagnostics.
	if d := tm26RedirectRouteExists(stml.ActionBlock{OperationID: "Login"}, "login.html", pages); len(d) != 0 {
		t.Errorf("no redirect: expected 0 diagnostics, got %+v", d)
	}
}
