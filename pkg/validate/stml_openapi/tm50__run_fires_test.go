//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what Run 경유 TM-50 — data-crumb-field 응답 필드 부재 ERROR 발화, sitemap 부재 시 침묵 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM50_RunFires(t *testing.T) {
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/buildings/{id}": getOp("GetBuilding", nil, map[string]*openapi3.SchemaRef{
			"building_name": stringProp(),
		}),
	})
	pages := []stml.PageSpec{
		{Name: "building-list", FileName: "building-list.html"},
		{Name: "building-detail", FileName: "building-detail.html", Fetches: []stml.FetchBlock{{OperationID: "GetBuilding"}}},
	}
	fs := makeFS(pages, doc)
	fs.Sitemap = &stml.SitemapSpec{
		FileName: "sitemap.html",
		Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{
			{Page: "building-list", Label: "건물 목록", Index: true, Menu: true, Children: []stml.SitemapNode{
				{Page: "building-detail", Label: "건물 상세", CrumbField: "no_such_field", Menu: true},
			}},
		}}},
	}

	if got := countDiag(Run(fs), "[TM-50]"); got != 1 {
		t.Errorf("expected 1 TM-50 via Run, got %d: %+v", got, Run(fs))
	}

	fs.Sitemap = nil
	if got := countDiag(Run(fs), "[TM-50]"); got != 0 {
		t.Errorf("expected 0 TM-50 via Run without a sitemap, got %d", got)
	}
}
