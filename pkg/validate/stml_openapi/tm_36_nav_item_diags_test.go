//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what tm36NavItemDiags — 단일 nav 항목 3 ERROR 분기와 합법 분기(인덱스/매칭/optional-only/패턴 없음) 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM36NavItemDiags(t *testing.T) {
	layout := stml.LayoutSpec{Name: "app", File: "layouts/app.html"}
	pages := []stml.PageSpec{
		{Name: "dashboard", FileName: "dashboard.html"},
		{Name: "building-detail", FileName: "building-detail.html",
			Fetches: []stml.FetchBlock{{
				OperationID: "GetBuilding",
				Params:      []stml.ParamBind{{Name: "building_id", Source: "route.BuildingID"}},
			}}},
	}

	t.Run("index path is silent", func(t *testing.T) {
		if diags := tm36NavItemDiags(layout, stml.NavItem{Path: "/"}, pages); len(diags) != 0 {
			t.Errorf("expected silence for \"/\", got %+v", diags)
		}
	})

	t.Run("matching static path is silent", func(t *testing.T) {
		if diags := tm36NavItemDiags(layout, stml.NavItem{Path: "/buildings/3"}, pages); len(diags) != 0 {
			t.Errorf("expected silence for concrete static path, got %+v", diags)
		}
	})

	t.Run("unmatched static path errors", func(t *testing.T) {
		diags := tm36NavItemDiags(layout, stml.NavItem{Path: "/nowhere"}, pages)
		if len(diags) != 1 || !strings.Contains(diags[0].Message, "does not resolve to any STML page route") {
			t.Errorf("expected static-path error, got %+v", diags)
		}
	})

	t.Run("known page without params is silent", func(t *testing.T) {
		if diags := tm36NavItemDiags(layout, stml.NavItem{Path: "dashboard"}, pages); len(diags) != 0 {
			t.Errorf("expected silence, got %+v", diags)
		}
	})

	t.Run("unknown page name errors", func(t *testing.T) {
		diags := tm36NavItemDiags(layout, stml.NavItem{Path: "nope"}, pages)
		if len(diags) != 1 || !strings.Contains(diags[0].Message, "does not name any STML page") {
			t.Errorf("expected page-name error, got %+v", diags)
		}
	})

	t.Run("required segment errors", func(t *testing.T) {
		diags := tm36NavItemDiags(layout, stml.NavItem{Path: "building-detail"}, pages)
		if len(diags) != 1 || !strings.Contains(diags[0].Message, "required segment :BuildingID") {
			t.Errorf("expected required-segment error, got %+v", diags)
		}
	})
}
