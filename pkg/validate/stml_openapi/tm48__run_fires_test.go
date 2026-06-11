//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what Run 경유 TM-48 — data-entry 블록 동적 그룹 ERROR 발화, sitemap 부재 시 침묵 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM48_RunFires(t *testing.T) {
	itemsProp := arrayProp("object")
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/buildings": getOp("ListMyBuildings", nil, map[string]*openapi3.SchemaRef{"items": itemsProp}),
	})
	pages := []stml.PageSpec{
		{Name: "building-detail", FileName: "building-detail.html", Fetches: []stml.FetchBlock{{OperationID: "ListMyBuildings"}}},
	}
	fs := makeFS(pages, doc)
	fs.Sitemap = &stml.SitemapSpec{
		FileName: "sitemap.html",
		Navs: []stml.SitemapNav{{Entry: true, Items: []stml.SitemapNode{
			{Page: "building-detail", Label: "건물 상세", Index: true, Menu: true},
			{Label: "내 건물", Menu: true, Fetch: "ListMyBuildings", Each: "items", Link: "building-detail", LabelField: "name"},
		}}},
	}

	if got := countDiag(Run(fs), "[TM-48]"); got != 1 {
		t.Errorf("expected 1 TM-48 via Run, got %d: %+v", got, Run(fs))
	}

	fs.Sitemap = nil
	if got := countDiag(Run(fs), "[TM-48]"); got != 0 {
		t.Errorf("expected 0 TM-48 via Run without a sitemap, got %d", got)
	}
}
