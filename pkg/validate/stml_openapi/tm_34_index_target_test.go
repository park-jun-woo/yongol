//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TM-34 — 없는 페이지명 / 필수 세그먼트 / "/" 동시 선언 3분기 발화와 합법 선언(보호·optional-only) 침묵 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM34IndexTarget(t *testing.T) {
	pagesOf := func(specs ...stml.PageSpec) []stml.PageSpec { return specs }

	t.Run("undeclared index is silent", func(t *testing.T) {
		fs := makeFS(pagesOf(stml.PageSpec{Name: "dashboard", FileName: "dashboard.html"}), nil)
		if diags := tm34IndexTarget(fs); len(diags) != 0 {
			t.Errorf("expected silence without frontend.index, got %+v", diags)
		}
	})

	t.Run("valid declaration is silent", func(t *testing.T) {
		fs := makeFS(pagesOf(
			stml.PageSpec{Name: "dashboard", FileName: "dashboard.html"},
			stml.PageSpec{Name: "login", FileName: "login.html"},
		), nil)
		fs.Manifest.Frontend.Index = "dashboard"
		if diags := tm34IndexTarget(fs); len(diags) != 0 {
			t.Errorf("expected silence for a valid index page, got %+v", diags)
		}
	})

	t.Run("optional-only segments are legal", func(t *testing.T) {
		// route.* consumed only by actions → ":Name?" — base path matches
		// with the segment omitted, so the page is a legal redirect target.
		fs := makeFS(pagesOf(stml.PageSpec{
			Name:     "unit-list",
			FileName: "unit-list.html",
			Actions: []stml.ActionBlock{{
				OperationID: "DeleteUnit",
				Params:      []stml.ParamBind{{Name: "building_id", Source: "route.BuildingID"}},
			}},
		}), nil)
		fs.Manifest.Frontend.Index = "unit-list"
		if diags := tm34IndexTarget(fs); len(diags) != 0 {
			t.Errorf("expected silence for optional-only segments, got %+v", diags)
		}
	})

	t.Run("unknown page name is an error", func(t *testing.T) {
		fs := makeFS(pagesOf(stml.PageSpec{Name: "dashboard", FileName: "dashboard.html"}), nil)
		fs.Manifest.Frontend.Index = "dashbord"
		diags := tm34IndexTarget(fs)
		if got := countDiag(diags, "[TM-34]"); got != 1 {
			t.Fatalf("expected 1 TM-34, got %d: %+v", got, diags)
		}
		if diags[0].Level != diagnostic.LevelError {
			t.Errorf("Level = %v, want LevelError", diags[0].Level)
		}
	})

	t.Run("required segment route is an error", func(t *testing.T) {
		fs := makeFS(pagesOf(stml.PageSpec{
			Name:     "building-detail",
			FileName: "building-detail.html",
			Fetches: []stml.FetchBlock{{
				OperationID: "GetBuilding",
				Params:      []stml.ParamBind{{Name: "id", Source: "route.BuildingID"}},
			}},
		}), nil)
		fs.Manifest.Frontend.Index = "building-detail"
		diags := tm34IndexTarget(fs)
		if got := countDiag(diags, "[TM-34]"); got != 1 {
			t.Fatalf("expected 1 TM-34 for required segment, got %d: %+v", got, diags)
		}
	})

	t.Run("simultaneous slash mount is an error", func(t *testing.T) {
		fs := makeFS(pagesOf(
			stml.PageSpec{Name: "home", FileName: "home.html", Route: "/"},
			stml.PageSpec{Name: "dashboard", FileName: "dashboard.html"},
		), nil)
		fs.Manifest.Frontend.Index = "dashboard"
		diags := tm34IndexTarget(fs)
		if got := countDiag(diags, "[TM-34]"); got != 1 {
			t.Fatalf("expected 1 TM-34 for the contradiction, got %d: %+v", got, diags)
		}
	})
}
