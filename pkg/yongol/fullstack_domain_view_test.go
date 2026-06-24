//ff:func feature=orchestrator type=test control=sequence
//ff:what TestDomainView — 단수 필드 5개 도메인 스왑 / 제약 공유 유지 / 부모 불변 검증

package yongol

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestDomainView(t *testing.T) {
	// Parent singular sentinels — must stay untouched after DomainView.
	parentDoc := &openapi3.T{}
	parentLines := &oapiparser.LineIndex{}
	parentPages := []stml.PageSpec{{Name: "parent"}}
	parentSitemap := &stml.SitemapSpec{FileName: "parent-sitemap.html"}
	parentLayouts := []stml.LayoutSpec{{Name: "parent-layout"}}

	// Per-domain data for "public".
	pubDoc := &openapi3.T{}
	pubLines := &oapiparser.LineIndex{}
	pubPages := []stml.PageSpec{{Name: "pub-page"}}
	pubSitemap := &stml.SitemapSpec{FileName: "pub-sitemap.html"}
	pubLayouts := []stml.LayoutSpec{{Name: "pub-layout"}}

	reqC := map[string]map[string]oapiparser.FieldConstraint{}
	respC := map[string]map[string]oapiparser.FieldConstraint{}

	fs := &Fullstack{
		OpenAPIDoc:          parentDoc,
		OpenAPILines:        parentLines,
		STMLPages:           parentPages,
		Sitemap:             parentSitemap,
		Layouts:             parentLayouts,
		RequestConstraints:  reqC,
		ResponseConstraints: respC,
		DomainOpenAPIDocs:   map[string]*openapi3.T{"public": pubDoc},
		DomainOpenAPILines:  map[string]*oapiparser.LineIndex{"public": pubLines},
		DomainSTMLPages:     map[string][]stml.PageSpec{"public": pubPages},
		DomainSitemaps:      map[string]*stml.SitemapSpec{"public": pubSitemap},
		DomainLayouts:       map[string][]stml.LayoutSpec{"public": pubLayouts},
	}

	view := fs.DomainView("public")

	// 1. The five singular fields are swapped to the domain's data.
	if view.OpenAPIDoc != pubDoc {
		t.Errorf("OpenAPIDoc = %p, want %p", view.OpenAPIDoc, pubDoc)
	}
	if view.OpenAPILines != pubLines {
		t.Errorf("OpenAPILines = %p, want %p", view.OpenAPILines, pubLines)
	}
	if view.Sitemap != pubSitemap {
		t.Errorf("Sitemap = %p, want %p", view.Sitemap, pubSitemap)
	}
	if len(view.STMLPages) != 1 || view.STMLPages[0].Name != "pub-page" {
		t.Errorf("STMLPages = %v, want [pub-page]", view.STMLPages)
	}
	if &view.STMLPages[0] != &pubPages[0] {
		t.Errorf("STMLPages backing array differs from domain's")
	}
	if len(view.Layouts) != 1 || view.Layouts[0].Name != "pub-layout" {
		t.Errorf("Layouts = %v, want [pub-layout]", view.Layouts)
	}
	if &view.Layouts[0] != &pubLayouts[0] {
		t.Errorf("Layouts backing array differs from domain's")
	}

	// 2. Constraints stay SHARED — same backing map as the parent. Maps are
	// reference types; a shallow copy shares the backing, so a write through one
	// is visible through the other.
	fs.RequestConstraints["sentinel-req"] = nil
	if _, ok := view.RequestConstraints["sentinel-req"]; !ok {
		t.Errorf("RequestConstraints not the shared parent map")
	}
	view.ResponseConstraints["sentinel-resp"] = nil
	if _, ok := fs.ResponseConstraints["sentinel-resp"]; !ok {
		t.Errorf("ResponseConstraints not the shared parent map")
	}

	// 3. The parent fs is not mutated (singular fields unchanged).
	if fs.OpenAPIDoc != parentDoc || fs.OpenAPILines != parentLines ||
		fs.Sitemap != parentSitemap {
		t.Errorf("parent singular pointer fields were mutated")
	}
	if len(fs.STMLPages) != 1 || fs.STMLPages[0].Name != "parent" {
		t.Errorf("parent STMLPages mutated = %v", fs.STMLPages)
	}
	if len(fs.Layouts) != 1 || fs.Layouts[0].Name != "parent-layout" {
		t.Errorf("parent Layouts mutated = %v", fs.Layouts)
	}
}
