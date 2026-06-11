//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TM-42 — data-index 중복(두 위치 표기) / page 없는 노드 발화와 단일·미선언 침묵 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM42SitemapIndexConflict(t *testing.T) {
	pages := []stml.PageSpec{
		{Name: "dashboard", FileName: "dashboard.html"},
		{Name: "login", FileName: "login.html"},
	}

	t.Run("valid single data-index is silent", func(t *testing.T) {
		fs := makeFS(pages, nil)
		fs.Sitemap = sitemapWithIndexFixture("dashboard")
		if diags := tm42SitemapIndexConflict(fs); len(diags) != 0 {
			t.Errorf("expected silence, got %+v", diags)
		}
	})

	t.Run("no data-index is silent", func(t *testing.T) {
		fs := makeFS(pages, nil)
		fs.Sitemap = &stml.SitemapSpec{
			FileName: "sitemap.html",
			Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{
				{Page: "dashboard", Label: "대시보드", Menu: true},
			}}},
		}
		if diags := tm42SitemapIndexConflict(fs); len(diags) != 0 {
			t.Errorf("expected silence, got %+v", diags)
		}
	})

	t.Run("two data-index entries fire with both positions", func(t *testing.T) {
		fs := makeFS(pages, nil)
		fs.Sitemap = &stml.SitemapSpec{
			FileName: "sitemap.html",
			Navs: []stml.SitemapNav{
				{Items: []stml.SitemapNode{{Page: "dashboard", Label: "대시보드", Index: true, Menu: true}}},
				{Items: []stml.SitemapNode{{Page: "login", Label: "로그인", Index: true, Menu: true}}},
			},
		}
		diags := tm42SitemapIndexConflict(fs)
		if got := countDiag(diags, "[TM-42]"); got != 1 {
			t.Fatalf("expected 1 TM-42, got %d: %+v", got, diags)
		}
		if diags[0].Level != diagnostic.LevelError {
			t.Errorf("Level = %v, want LevelError", diags[0].Level)
		}
		if !strings.Contains(diags[0].Message, "nav[0] > 대시보드") || !strings.Contains(diags[0].Message, "nav[1] > 로그인") {
			t.Errorf("Message should list both positions, got %q", diags[0].Message)
		}
	})

	t.Run("data-index without data-page fires", func(t *testing.T) {
		fs := makeFS(pages, nil)
		fs.Sitemap = &stml.SitemapSpec{
			FileName: "sitemap.html",
			Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{
				{Label: "그룹", Index: true, Menu: true},
			}}},
		}
		diags := tm42SitemapIndexConflict(fs)
		if got := countDiag(diags, "[TM-42]"); got != 1 {
			t.Fatalf("expected 1 TM-42, got %d: %+v", got, diags)
		}
		if !strings.Contains(diags[0].Message, "without data-page") {
			t.Errorf("Message should name the missing data-page, got %q", diags[0].Message)
		}
	})
}
