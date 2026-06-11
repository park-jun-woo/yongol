//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TM-30 사이트맵 확장 — data-label-field 누락/item 스키마 미존재/비라벨 타입 ERROR, 정상·미해석 침묵 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM30SitemapGroupLabelField(t *testing.T) {
	itemsProp := &openapi3.SchemaRef{Value: &openapi3.Schema{
		Type: &openapi3.Types{"array"},
		Items: &openapi3.SchemaRef{Value: &openapi3.Schema{
			Type: &openapi3.Types{"object"},
			Properties: map[string]*openapi3.SchemaRef{
				"building_id":   intProp(),
				"building_name": stringProp(),
				"meta":          objectProp(),
			},
		}},
	}}
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/buildings": getOp("ListMyBuildings", nil, map[string]*openapi3.SchemaRef{"items": itemsProp}),
	})
	raif := map[string]map[string]map[string]bool{
		"ListMyBuildings": {"items": {"building_id": true, "building_name": true, "meta": true}},
	}
	group := func(label string) *stml.SitemapSpec {
		return &stml.SitemapSpec{FileName: "sitemap.html", Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{
			{Label: "내 건물", Fetch: "ListMyBuildings", Each: "items", Link: "building-detail", LabelField: label},
		}}}}
	}

	t.Run("string label field is silent", func(t *testing.T) {
		fs := makeFS(nil, doc)
		fs.Sitemap = group("building_name")
		if diags := tm30SitemapGroupLabelField(fs, raif); len(diags) != 0 {
			t.Errorf("expected silence, got %+v", diags)
		}
	})

	t.Run("missing data-label-field is an ERROR — the attribute is required", func(t *testing.T) {
		fs := makeFS(nil, doc)
		fs.Sitemap = group("")
		diags := tm30SitemapGroupLabelField(fs, raif)
		if got := countDiag(diags, "[TM-30]"); got != 1 {
			t.Fatalf("expected 1 TM-30, got %d: %+v", got, diags)
		}
		if !strings.Contains(diags[0].Message, "requires data-label-field") {
			t.Errorf("message = %q", diags[0].Message)
		}
	})

	t.Run("field absent from the item schema is an ERROR", func(t *testing.T) {
		fs := makeFS(nil, doc)
		fs.Sitemap = group("nope")
		diags := tm30SitemapGroupLabelField(fs, raif)
		if got := countDiag(diags, "[TM-30]"); got != 1 {
			t.Fatalf("expected 1 TM-30, got %d: %+v", got, diags)
		}
		if !strings.Contains(diags[0].Message, "not in the item schema") {
			t.Errorf("message = %q", diags[0].Message)
		}
	})

	t.Run("non-scalar field type is an ERROR", func(t *testing.T) {
		fs := makeFS(nil, doc)
		fs.Sitemap = group("meta")
		diags := tm30SitemapGroupLabelField(fs, raif)
		if got := countDiag(diags, "[TM-30]"); got != 1 {
			t.Fatalf("expected 1 TM-30, got %d: %+v", got, diags)
		}
		if !strings.Contains(diags[0].Message, `"object"`) {
			t.Errorf("message should name the offending type, got %q", diags[0].Message)
		}
	})

	t.Run("integer label field is legal", func(t *testing.T) {
		fs := makeFS(nil, doc)
		fs.Sitemap = group("building_id")
		if diags := tm30SitemapGroupLabelField(fs, raif); len(diags) != 0 {
			t.Errorf("expected silence, got %+v", diags)
		}
	})

	t.Run("unresolvable item schema stays silent (TM-01/07 own it)", func(t *testing.T) {
		fs := makeFS(nil, doc)
		fs.Sitemap = &stml.SitemapSpec{FileName: "sitemap.html", Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{
			{Label: "내 건물", Fetch: "Nope", Each: "items", Link: "building-detail", LabelField: "building_name"},
		}}}}
		if diags := tm30SitemapGroupLabelField(fs, raif); len(diags) != 0 {
			t.Errorf("expected silence, got %+v", diags)
		}
	})
}
