//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TM-39 — 미실존 data-page ERROR / data-page+href 동시 보유 ERROR / 그룹 li 의 data-crumb-field ERROR / 정상·그룹·외부 링크 침묵 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM39SitemapPageNotFound(t *testing.T) {
	pages := []stml.PageSpec{
		{Name: "login", FileName: "login.html"},
		{Name: "dashboard", FileName: "dashboard.html"},
	}

	t.Run("nonexistent page fires with position", func(t *testing.T) {
		fs := makeFS(pages, nil)
		fs.Sitemap = &stml.SitemapSpec{
			FileName: "sitemap.html",
			Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{
				{Label: "관리", Children: []stml.SitemapNode{
					{Page: "no-such-page", Label: "유령", Menu: true},
				}},
			}}},
		}
		diags := tm39SitemapPageNotFound(fs)
		if got := countDiag(diags, "[TM-39]"); got != 1 {
			t.Fatalf("expected 1 TM-39, got %d: %+v", got, diags)
		}
		if diags[0].Level != diagnostic.LevelError {
			t.Errorf("Level = %v, want LevelError", diags[0].Level)
		}
		if !strings.Contains(diags[0].Message, "no-such-page") || !strings.Contains(diags[0].Message, "nav[0] > 관리 > 유령") {
			t.Errorf("Message should name the page and its position, got %q", diags[0].Message)
		}
	})

	t.Run("data-page and href together fires", func(t *testing.T) {
		fs := makeFS(pages, nil)
		fs.Sitemap = &stml.SitemapSpec{
			FileName: "sitemap.html",
			Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{
				{Page: "login", Href: "https://example.com", Label: "로그인", Menu: true},
			}}},
		}
		diags := tm39SitemapPageNotFound(fs)
		if got := countDiag(diags, "[TM-39]"); got != 1 {
			t.Fatalf("expected 1 TM-39 for the page+href conflict, got %d: %+v", got, diags)
		}
		if !strings.Contains(diags[0].Message, "mutually exclusive") {
			t.Errorf("Message should state the mutual exclusion, got %q", diags[0].Message)
		}
	})

	t.Run("data-crumb-field on a page-less group li fires (Phase006)", func(t *testing.T) {
		fs := makeFS(pages, nil)
		fs.Sitemap = &stml.SitemapSpec{
			FileName: "sitemap.html",
			Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{
				{Label: "건물", CrumbField: "building_name", Menu: true, Children: []stml.SitemapNode{
					{Page: "dashboard", Label: "대시보드", Menu: true},
				}},
			}}},
		}
		diags := tm39SitemapPageNotFound(fs)
		if got := countDiag(diags, "[TM-39]"); got != 1 {
			t.Fatalf("expected 1 TM-39 for the misplaced data-crumb-field, got %d: %+v", got, diags)
		}
		if !strings.Contains(diags[0].Message, "data-crumb-field") || !strings.Contains(diags[0].Message, "page items only") {
			t.Errorf("Message should state the page-item-only placement, got %q", diags[0].Message)
		}
	})

	t.Run("data-crumb-field on a page item is silent here (TM-50's concern)", func(t *testing.T) {
		fs := makeFS(pages, nil)
		fs.Sitemap = &stml.SitemapSpec{
			FileName: "sitemap.html",
			Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{
				{Page: "dashboard", Label: "대시보드", CrumbField: "name", Menu: true},
			}}},
		}
		if diags := tm39SitemapPageNotFound(fs); len(diags) != 0 {
			t.Errorf("expected silence, got %+v", diags)
		}
	})

	t.Run("existing pages, groups and external links are silent", func(t *testing.T) {
		fs := makeFS(pages, nil)
		fs.Sitemap = &stml.SitemapSpec{
			FileName: "sitemap.html",
			Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{
				{Page: "dashboard", Label: "대시보드", Menu: true},
				{Label: "그룹", Menu: true, Children: []stml.SitemapNode{
					{Page: "login", Label: "로그인", Menu: true},
				}},
				{Href: "https://docs.example.com", Label: "문서", Menu: true},
			}}},
		}
		if diags := tm39SitemapPageNotFound(fs); len(diags) != 0 {
			t.Errorf("expected silence, got %+v", diags)
		}
	})
}
