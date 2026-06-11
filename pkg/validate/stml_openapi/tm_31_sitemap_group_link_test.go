//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TM-31 사이트맵 확장 — 동적 그룹 data-link 대상 페이지 미존재 ERROR / 실존·link 부재 침묵 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM31SitemapGroupLink(t *testing.T) {
	pages := []stml.PageSpec{{Name: "building-detail", FileName: "building-detail.html"}}
	group := func(link string) *stml.SitemapSpec {
		return &stml.SitemapSpec{FileName: "sitemap.html", Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{
			{Label: "내 건물", Fetch: "ListMyBuildings", Each: "items", Link: link, LabelField: "name"},
		}}}}
	}

	t.Run("existing target page is silent", func(t *testing.T) {
		fs := makeFS(pages, nil)
		fs.Sitemap = group("building-detail")
		if diags := tm31SitemapGroupLink(fs); len(diags) != 0 {
			t.Errorf("expected silence, got %+v", diags)
		}
	})

	t.Run("unknown target page fires", func(t *testing.T) {
		fs := makeFS(pages, nil)
		fs.Sitemap = group("nope-detail")
		diags := tm31SitemapGroupLink(fs)
		if got := countDiag(diags, "[TM-31]"); got != 1 {
			t.Fatalf("expected 1 TM-31, got %d: %+v", got, diags)
		}
		if !strings.Contains(diags[0].Message, `"nope-detail"`) {
			t.Errorf("message = %q", diags[0].Message)
		}
	})

	t.Run("a group without data-link is TM-48's finding", func(t *testing.T) {
		fs := makeFS(pages, nil)
		fs.Sitemap = group("")
		if diags := tm31SitemapGroupLink(fs); len(diags) != 0 {
			t.Errorf("expected silence, got %+v", diags)
		}
	})
}
