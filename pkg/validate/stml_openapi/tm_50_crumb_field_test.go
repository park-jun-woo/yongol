//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TM-50 — fetch 부재(a)/응답 필드 없음(b)/라벨 불가 타입(c) ERROR, string·integer 정상·미실존 페이지·미지 op 침묵 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestTM50CrumbField(t *testing.T) {
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/buildings/{id}": getOp("GetBuilding", nil, map[string]*openapi3.SchemaRef{
			"building_name": stringProp(),
			"floor_count":   intProp(),
			"owner":         objectProp(),
			"tags":          arrayProp("string"),
		}),
	})
	opMap := buildOperationMethodMap(doc)
	withSitemap := func(pages []stml.PageSpec, field string) *yongol.Fullstack {
		fs := makeFS(pages, doc)
		fs.Sitemap = &stml.SitemapSpec{
			FileName: "sitemap.html",
			Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{
				{Page: "building-list", Label: "건물 목록", Menu: true, Children: []stml.SitemapNode{
					{Page: "building-detail", Label: "건물 상세", CrumbField: field, Menu: true},
				}},
			}}},
		}
		return fs
	}
	fetchPages := []stml.PageSpec{
		{Name: "building-list", FileName: "building-list.html"},
		{Name: "building-detail", FileName: "building-detail.html", Fetches: []stml.FetchBlock{{OperationID: "GetBuilding"}}},
	}

	t.Run("(a) page without a data-fetch fires", func(t *testing.T) {
		pages := []stml.PageSpec{
			{Name: "building-list", FileName: "building-list.html"},
			{Name: "building-detail", FileName: "building-detail.html"},
		}
		diags := tm50CrumbField(withSitemap(pages, "building_name"), opMap)
		if got := countDiag(diags, "[TM-50]"); got != 1 {
			t.Fatalf("expected 1 TM-50, got %d: %+v", got, diags)
		}
		if diags[0].Level != diagnostic.LevelError {
			t.Errorf("Level = %v, want LevelError", diags[0].Level)
		}
		if !strings.Contains(diags[0].Message, "needs a data-fetch") {
			t.Errorf("Message should name the missing fetch, got %q", diags[0].Message)
		}
	})

	t.Run("(b) field absent from the first fetch's 2xx response fires", func(t *testing.T) {
		diags := tm50CrumbField(withSitemap(fetchPages, "no_such_field"), opMap)
		if got := countDiag(diags, "[TM-50]"); got != 1 {
			t.Fatalf("expected 1 TM-50, got %d: %+v", got, diags)
		}
		if !strings.Contains(diags[0].Message, "no_such_field") || !strings.Contains(diags[0].Message, "GetBuilding") {
			t.Errorf("Message should name the field and the operation, got %q", diags[0].Message)
		}
	})

	t.Run("(c) object and array fields fire — not label-renderable", func(t *testing.T) {
		for _, field := range []string{"owner", "tags"} {
			diags := tm50CrumbField(withSitemap(fetchPages, field), opMap)
			if got := countDiag(diags, "[TM-50]"); got != 1 {
				t.Fatalf("expected 1 TM-50 for %q, got %d: %+v", field, got, diags)
			}
			if !strings.Contains(diags[0].Message, "crumb label") {
				t.Errorf("Message should state the label constraint, got %q", diags[0].Message)
			}
		}
	})

	t.Run("string and integer fields pass", func(t *testing.T) {
		for _, field := range []string{"building_name", "floor_count"} {
			if diags := tm50CrumbField(withSitemap(fetchPages, field), opMap); len(diags) != 0 {
				t.Errorf("expected silence for %q, got %+v", field, diags)
			}
		}
	})

	t.Run("unknown page and unknown operationId are other rules' findings", func(t *testing.T) {
		ghost := withSitemap(nil, "building_name") // no STML pages — TM-39 reports
		if diags := tm50CrumbField(ghost, opMap); len(diags) != 0 {
			t.Errorf("expected silence for the unknown page, got %+v", diags)
		}
		unknownOp := []stml.PageSpec{
			{Name: "building-list", FileName: "building-list.html"},
			{Name: "building-detail", FileName: "building-detail.html", Fetches: []stml.FetchBlock{{OperationID: "NoSuchOp"}}},
		}
		if diags := tm50CrumbField(withSitemap(unknownOp, "building_name"), opMap); len(diags) != 0 {
			t.Errorf("expected silence for the unknown operationId, got %+v", diags)
		}
	})
}
