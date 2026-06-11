//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what Run 경유 TM-43 — sitemap 존재 + 고아 페이지 WARNING 발화, sitemap 부재 시 침묵 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM43_RunFires(t *testing.T) {
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/buildings": getOp("ListBuildings", nil, nil),
	})
	pages := []stml.PageSpec{
		{Name: "building-list", FileName: "building-list.html", Fetches: []stml.FetchBlock{{OperationID: "ListBuildings"}}},
		{Name: "building-detail", FileName: "building-detail.html", Route: "/buildings/:BuildingID"},
	}
	fs := makeFS(pages, doc)
	fs.Sitemap = &stml.SitemapSpec{
		FileName: "sitemap.html",
		Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{
			{Page: "building-list", Label: "건물 목록", Index: true, Menu: true, Children: []stml.SitemapNode{
				{Page: "building-detail", Label: "건물 상세", Menu: true},
			}},
		}}},
	}

	diags := Run(fs)
	if got := countDiag(diags, "[TM-43]"); got != 1 {
		t.Errorf("expected 1 TM-43 via Run (listed detail page without an incoming link), got %d: %+v", got, diags)
	}

	fs.Sitemap = nil
	diags = Run(fs)
	if got := countDiag(diags, "[TM-43]"); got != 0 {
		t.Errorf("expected 0 TM-43 via Run without a sitemap, got %d: %+v", got, diags)
	}
}
