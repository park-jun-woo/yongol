//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TM-07/08 사이트맵 확장 — 동적 그룹 data-each 응답 부재/비배열 ERROR, 정상·미해석 op 침묵 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM07SitemapGroupEach(t *testing.T) {
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/buildings": getOp("ListMyBuildings", nil, map[string]*openapi3.SchemaRef{
			"items": arrayProp("object"),
			"total": intProp(),
		}),
	})
	opMap := buildOperationMethodMap(doc)
	group := func(fetch, each string) *stml.SitemapSpec {
		return &stml.SitemapSpec{FileName: "sitemap.html", Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{
			{Label: "내 건물", Fetch: fetch, Each: each, Link: "building-detail", LabelField: "name"},
		}}}}
	}

	t.Run("array field is silent", func(t *testing.T) {
		fs := makeFS(nil, doc)
		fs.Sitemap = group("ListMyBuildings", "items")
		if diags := tm07SitemapGroupEach(fs, opMap); len(diags) != 0 {
			t.Errorf("expected silence, got %+v", diags)
		}
	})

	t.Run("field missing from the response fires TM-07", func(t *testing.T) {
		fs := makeFS(nil, doc)
		fs.Sitemap = group("ListMyBuildings", "rows")
		diags := tm07SitemapGroupEach(fs, opMap)
		if got := countDiag(diags, "[TM-07]"); got != 1 {
			t.Fatalf("expected 1 TM-07, got %d: %+v", got, diags)
		}
	})

	t.Run("non-array field fires TM-08", func(t *testing.T) {
		fs := makeFS(nil, doc)
		fs.Sitemap = group("ListMyBuildings", "total")
		diags := tm07SitemapGroupEach(fs, opMap)
		if got := countDiag(diags, "[TM-08]"); got != 1 {
			t.Fatalf("expected 1 TM-08, got %d: %+v", got, diags)
		}
	})

	t.Run("unknown op stays silent (TM-01 owns it)", func(t *testing.T) {
		fs := makeFS(nil, doc)
		fs.Sitemap = group("Nope", "items")
		if diags := tm07SitemapGroupEach(fs, opMap); len(diags) != 0 {
			t.Errorf("expected silence, got %+v", diags)
		}
	})
}
