//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TM-48 — data-entry 블록의 동적 그룹 ERROR / 필수 어휘 누락 ERROR / 완전·정적은 침묵 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM48SitemapDynamicGroup(t *testing.T) {
	complete := stml.SitemapNode{Label: "내 건물", Fetch: "ListMyBuildings", Each: "items", Link: "building-detail", LabelField: "building_name"}

	t.Run("a complete group in a signed-in block is silent", func(t *testing.T) {
		fs := makeFS(nil, nil)
		fs.Sitemap = &stml.SitemapSpec{FileName: "sitemap.html", Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{complete}}}}
		if diags := tm48SitemapDynamicGroup(fs); len(diags) != 0 {
			t.Errorf("expected silence, got %+v", diags)
		}
	})

	t.Run("a dynamic group in a data-entry block is an ERROR", func(t *testing.T) {
		fs := makeFS(nil, nil)
		fs.Sitemap = &stml.SitemapSpec{FileName: "sitemap.html", Navs: []stml.SitemapNav{{Entry: true, Items: []stml.SitemapNode{complete}}}}
		diags := tm48SitemapDynamicGroup(fs)
		if got := countDiag(diags, "[TM-48]"); got != 1 {
			t.Fatalf("expected 1 TM-48, got %d: %+v", got, diags)
		}
		if diags[0].Level != diagnostic.LevelError || !strings.Contains(diags[0].Message, "data-entry") {
			t.Errorf("diag = %+v", diags[0])
		}
	})

	t.Run("missing required vocabulary is an ERROR naming the gaps", func(t *testing.T) {
		fs := makeFS(nil, nil)
		fs.Sitemap = &stml.SitemapSpec{FileName: "sitemap.html", Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{
			{Label: "내 건물", Each: "items", LabelField: "building_name"},
		}}}}
		diags := tm48SitemapDynamicGroup(fs)
		if got := countDiag(diags, "[TM-48]"); got != 1 {
			t.Fatalf("expected 1 TM-48, got %d: %+v", got, diags)
		}
		if !strings.Contains(diags[0].Message, "data-fetch") || !strings.Contains(diags[0].Message, "data-link") {
			t.Errorf("message should name the missing attributes, got %q", diags[0].Message)
		}
		// a missing data-label-field alone is TM-30's finding, not TM-48's
		if strings.Contains(diags[0].Message, "missing data-label-field") {
			t.Errorf("label-field absence belongs to TM-30, got %q", diags[0].Message)
		}
	})

	t.Run("missing data-each alone is an ERROR naming only that gap", func(t *testing.T) {
		fs := makeFS(nil, nil)
		fs.Sitemap = &stml.SitemapSpec{FileName: "sitemap.html", Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{
			{Label: "내 건물", Fetch: "ListMyBuildings", Link: "building-detail"},
		}}}}
		diags := tm48SitemapDynamicGroup(fs)
		if got := countDiag(diags, "[TM-48]"); got != 1 {
			t.Fatalf("expected 1 TM-48, got %d: %+v", got, diags)
		}
		if !strings.Contains(diags[0].Message, "data-each") {
			t.Errorf("message should name data-each as missing, got %q", diags[0].Message)
		}
		if strings.Contains(diags[0].Message, "missing data-fetch") || strings.Contains(diags[0].Message, "data-link,") {
			t.Errorf("only data-each should be reported missing, got %q", diags[0].Message)
		}
	})

	t.Run("static nodes are silent", func(t *testing.T) {
		fs := makeFS(nil, nil)
		fs.Sitemap = &stml.SitemapSpec{FileName: "sitemap.html", Navs: []stml.SitemapNav{{Entry: true, Items: []stml.SitemapNode{{Page: "login", Label: "로그인"}}}}}
		if diags := tm48SitemapDynamicGroup(fs); len(diags) != 0 {
			t.Errorf("expected silence, got %+v", diags)
		}
	})
}
