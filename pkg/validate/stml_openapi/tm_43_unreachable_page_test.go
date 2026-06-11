//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TM-43 — 등재≠도달 헌법 케이스(등재+필수 파라미터+무링크 → WARNING) / 메뉴·link 도달 침묵 / entry 루트 / sitemap 부재 침묵 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM43UnreachablePage(t *testing.T) {
	detailPage := stml.PageSpec{Name: "building-detail", FileName: "building-detail.html", Route: "/buildings/:BuildingID"}
	listPage := stml.PageSpec{Name: "building-list", FileName: "building-list.html"}
	sitemapBoth := func() *stml.SitemapSpec {
		return &stml.SitemapSpec{
			FileName: "sitemap.html",
			Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{
				{Page: "building-list", Label: "건물 목록", Index: true, Menu: true, Children: []stml.SitemapNode{
					{Page: "building-detail", Label: "건물 상세", Menu: true},
				}},
			}}},
		}
	}

	t.Run("listed but no required-param fill and no incoming link fires — listing is not reaching", func(t *testing.T) {
		// The constitutional case of plans/stml/sitemap Phase002 (DESIGN
		// §4.10, BUG-122): building-detail IS in the sitemap, yet its route
		// needs :BuildingID, no menu renders it and nothing links to it.
		fs := makeFS([]stml.PageSpec{listPage, detailPage}, nil)
		fs.Sitemap = sitemapBoth()
		diags := tm43UnreachablePage(fs)
		if got := countDiag(diags, "[TM-43]"); got != 1 {
			t.Fatalf("expected 1 TM-43 despite the sitemap listing, got %d: %+v", got, diags)
		}
		if diags[0].Level != diagnostic.LevelWarning {
			t.Errorf("Level = %v, want LevelWarning", diags[0].Level)
		}
		if !strings.Contains(diags[0].Message, "building-detail") ||
			!strings.Contains(diags[0].Message, "in sitemap but not menu-rendered") ||
			!strings.Contains(diags[0].Message, "no data-link/data-redirect/breadcrumb edge points to it") {
			t.Errorf("Message should classify the cause, got %q", diags[0].Message)
		}
	})

	t.Run("a data-link from the menu-rendered list silences it", func(t *testing.T) {
		linked := listPage
		linked.Children = []stml.ChildNode{
			{Kind: "each", Each: &stml.EachBlock{Field: "items", RowLink: &stml.LinkRef{TargetPage: "building-detail"}}},
		}
		fs := makeFS([]stml.PageSpec{linked, detailPage}, nil)
		fs.Sitemap = sitemapBoth()
		if diags := tm43UnreachablePage(fs); len(diags) != 0 {
			t.Errorf("expected silence once the list row links to the detail, got %+v", diags)
		}
	})

	t.Run("menu-rendered entries alone reach", func(t *testing.T) {
		fs := makeFS([]stml.PageSpec{listPage, {Name: "settings", FileName: "settings.html"}}, nil)
		fs.Sitemap = &stml.SitemapSpec{
			FileName: "sitemap.html",
			Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{
				{Page: "building-list", Label: "건물", Index: true, Menu: true},
				{Page: "settings", Label: "설정", Menu: true},
			}}},
		}
		if diags := tm43UnreachablePage(fs); len(diags) != 0 {
			t.Errorf("expected silence for menu-rendered pages, got %+v", diags)
		}
	})

	t.Run("a data-entry block makes its pages roots", func(t *testing.T) {
		fs := makeFS([]stml.PageSpec{{Name: "login", FileName: "login.html"}}, nil)
		fs.Sitemap = &stml.SitemapSpec{
			FileName: "sitemap.html",
			Navs: []stml.SitemapNav{{Entry: true, Items: []stml.SitemapNode{
				{Page: "login", Label: "로그인", Menu: false},
			}}},
		}
		if diags := tm43UnreachablePage(fs); len(diags) != 0 {
			t.Errorf("expected silence for an entry-block page, got %+v", diags)
		}
	})

	t.Run("completely unlisted orphan names the other cause", func(t *testing.T) {
		fs := makeFS([]stml.PageSpec{listPage, {Name: "lost", FileName: "lost.html"}}, nil)
		fs.Sitemap = &stml.SitemapSpec{
			FileName: "sitemap.html",
			Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{
				{Page: "building-list", Label: "건물", Index: true, Menu: true},
			}}},
		}
		diags := tm43UnreachablePage(fs)
		if got := countDiag(diags, "[TM-43]"); got != 1 {
			t.Fatalf("expected 1 TM-43 for the unlisted page, got %d: %+v", got, diags)
		}
		if !strings.Contains(diags[0].Message, "not in the sitemap") {
			t.Errorf("Message should say the page is not in the sitemap, got %q", diags[0].Message)
		}
	})

	t.Run("sitemap absent stays silent", func(t *testing.T) {
		fs := makeFS([]stml.PageSpec{listPage, detailPage}, nil)
		if diags := tm43UnreachablePage(fs); len(diags) != 0 {
			t.Errorf("expected silence without a sitemap (TM-49's territory), got %+v", diags)
		}
	})
}
