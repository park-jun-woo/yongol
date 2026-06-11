//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TM-01 사이트맵 확장 — 동적 그룹 data-fetch 미존재 op ERROR / 실존 op·fetch 부재 침묵 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM01SitemapGroupFetch(t *testing.T) {
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/buildings": getOp("ListMyBuildings", nil, map[string]*openapi3.SchemaRef{"items": arrayProp("object")}),
	})
	opMap := buildOperationMethodMap(doc)
	group := func(fetch string) *stml.SitemapSpec {
		return &stml.SitemapSpec{FileName: "sitemap.html", Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{
			{Label: "내 건물", Fetch: fetch, Each: "items", Link: "building-detail", LabelField: "name"},
		}}}}
	}

	t.Run("unknown operationId fires", func(t *testing.T) {
		fs := makeFS(nil, doc)
		fs.Sitemap = group("Nope")
		diags := tm01SitemapGroupFetch(fs, opMap)
		if got := countDiag(diags, "[TM-01]"); got != 1 {
			t.Fatalf("expected 1 TM-01, got %d: %+v", got, diags)
		}
		if !strings.Contains(diags[0].Message, `"Nope"`) || !strings.Contains(diags[0].Message, "내 건물") {
			t.Errorf("message = %q", diags[0].Message)
		}
	})

	t.Run("existing operationId is silent", func(t *testing.T) {
		fs := makeFS(nil, doc)
		fs.Sitemap = group("ListMyBuildings")
		if diags := tm01SitemapGroupFetch(fs, opMap); len(diags) != 0 {
			t.Errorf("expected silence, got %+v", diags)
		}
	})

	t.Run("a group without data-fetch is TM-48's finding, not TM-01's", func(t *testing.T) {
		fs := makeFS(nil, doc)
		fs.Sitemap = group("")
		if diags := tm01SitemapGroupFetch(fs, opMap); len(diags) != 0 {
			t.Errorf("expected silence, got %+v", diags)
		}
	})
}
