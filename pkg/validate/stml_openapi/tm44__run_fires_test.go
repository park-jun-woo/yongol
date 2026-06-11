//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what Run 경유 TM-44 — sitemap 존재 + data-nav 잔존 ERROR 발화, sitemap 부재 시 침묵 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM44_RunFires(t *testing.T) {
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/dashboard": getOp("GetDashboard", nil, nil),
	})
	pages := []stml.PageSpec{
		{Name: "dashboard", FileName: "dashboard.html", Fetches: []stml.FetchBlock{{OperationID: "GetDashboard"}}},
	}
	fs := makeFS(pages, doc)
	fs.Layouts = []stml.LayoutSpec{{
		Name: "app", File: "layouts/app.html",
		NavItems: []stml.NavItem{{Path: "dashboard", Label: "대시보드"}},
	}}
	fs.Sitemap = &stml.SitemapSpec{
		FileName: "sitemap.html",
		Navs: []stml.SitemapNav{{Layout: "app", Items: []stml.SitemapNode{
			{Page: "dashboard", Label: "대시보드", Index: true, Menu: true},
		}}},
	}

	diags := Run(fs)
	if got := countDiag(diags, "[TM-44]"); got != 1 {
		t.Errorf("expected 1 TM-44 via Run (data-nav surviving next to the sitemap), got %d: %+v", got, diags)
	}

	fs.Sitemap = nil
	diags = Run(fs)
	if got := countDiag(diags, "[TM-44]"); got != 0 {
		t.Errorf("expected 0 TM-44 via Run without a sitemap, got %d: %+v", got, diags)
	}
}
