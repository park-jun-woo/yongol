//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TM-36 — 페이지명/정적 경로/인덱스/오타/미매칭/필수 세그먼트/optional-only 매트릭스 발화·침묵 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM36NavTarget(t *testing.T) {
	layoutOf := func(navs ...stml.NavItem) []stml.LayoutSpec {
		return []stml.LayoutSpec{{Name: "app", File: "layouts/app.html", NavItems: navs}}
	}
	pages := []stml.PageSpec{
		{Name: "dashboard", FileName: "dashboard.html"},
		{Name: "building-list", FileName: "building-list.html"},
		{Name: "building-detail", FileName: "building-detail.html",
			Fetches: []stml.FetchBlock{{
				OperationID: "GetBuilding",
				Params:      []stml.ParamBind{{Name: "building_id", Source: "route.BuildingID"}},
			}}},
		{Name: "unit-list", FileName: "unit-list.html",
			Actions: []stml.ActionBlock{{
				OperationID: "DeleteUnit",
				Params:      []stml.ParamBind{{Name: "unit_id", Source: "route.UnitID"}},
			}}},
	}

	t.Run("page-name reference is silent", func(t *testing.T) {
		diags := tm36NavTarget(layoutOf(stml.NavItem{Path: "building-list", Label: "건물"}), pages)
		if len(diags) != 0 {
			t.Errorf("expected silence, got %+v", diags)
		}
	})

	t.Run("static path matching a route is silent", func(t *testing.T) {
		diags := tm36NavTarget(layoutOf(stml.NavItem{Path: "/dashboard", Label: "대시보드"}), pages)
		if len(diags) != 0 {
			t.Errorf("expected silence, got %+v", diags)
		}
	})

	t.Run("root static path is silent", func(t *testing.T) {
		diags := tm36NavTarget(layoutOf(stml.NavItem{Path: "/", Label: "홈"}), pages)
		if len(diags) != 0 {
			t.Errorf("expected silence for index path, got %+v", diags)
		}
	})

	t.Run("optional-only segments are silent", func(t *testing.T) {
		// route.UnitID is action-only → ":UnitID?" — the emitter strips it.
		diags := tm36NavTarget(layoutOf(stml.NavItem{Path: "unit-list", Label: "세대"}), pages)
		if len(diags) != 0 {
			t.Errorf("expected silence for optional-only route, got %+v", diags)
		}
	})

	t.Run("typo page name is an error", func(t *testing.T) {
		diags := tm36NavTarget(layoutOf(stml.NavItem{Path: "building-detial", Label: "건물"}), pages)
		if got := countDiag(diags, "[TM-36]"); got != 1 {
			t.Fatalf("expected 1 TM-36, got %d: %+v", got, diags)
		}
		if diags[0].Level != diagnostic.LevelError {
			t.Errorf("Level = %v, want LevelError", diags[0].Level)
		}
		if diags[0].File != "layouts/app.html" {
			t.Errorf("File = %q, want layout file", diags[0].File)
		}
		if !strings.Contains(diags[0].Message, "does not name any STML page") {
			t.Errorf("Message = %q, want page-name branch", diags[0].Message)
		}
	})

	t.Run("unmatched static path is an error", func(t *testing.T) {
		diags := tm36NavTarget(layoutOf(stml.NavItem{Path: "/nowhere", Label: "?"}), pages)
		if got := countDiag(diags, "[TM-36]"); got != 1 {
			t.Fatalf("expected 1 TM-36, got %d: %+v", got, diags)
		}
		if !strings.Contains(diags[0].Message, "does not resolve to any STML page route") {
			t.Errorf("Message = %q, want static-path branch", diags[0].Message)
		}
	})

	t.Run("required segment target is an error", func(t *testing.T) {
		// building-detail derives /buildings/:BuildingID — fetch-consumed,
		// required. A static menu link has no value to fill it.
		diags := tm36NavTarget(layoutOf(stml.NavItem{Path: "building-detail", Label: "건물 상세"}), pages)
		if got := countDiag(diags, "[TM-36]"); got != 1 {
			t.Fatalf("expected 1 TM-36, got %d: %+v", got, diags)
		}
		if !strings.Contains(diags[0].Message, "required segment :BuildingID") {
			t.Errorf("Message = %q, want required-segment branch", diags[0].Message)
		}
	})

	t.Run("no layouts is silent", func(t *testing.T) {
		if diags := tm36NavTarget(nil, pages); len(diags) != 0 {
			t.Errorf("expected silence without layouts, got %+v", diags)
		}
	})
}
