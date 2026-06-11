//ff:func feature=gen-react type=test control=iteration dimension=1
//ff:what sitemapPageLayouts — 중첩 페이지 포함 블록 레이아웃 매핑/"" 블록 제외/nil sitemap 검증

package react

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestSitemapPageLayouts(t *testing.T) {
	sitemap := &stml.SitemapSpec{Navs: []stml.SitemapNav{
		{Layout: "app", Items: []stml.SitemapNode{
			{Page: "dashboard", Menu: true},
			{Label: "그룹", Menu: true, Children: []stml.SitemapNode{
				{Page: "building-list", Menu: true, Children: []stml.SitemapNode{
					{Page: "building-detail", Menu: true},
				}},
			}},
		}},
		{Layout: "", Items: []stml.SitemapNode{{Page: "login", Menu: true}}},
	}}

	got := sitemapPageLayouts(sitemap)
	for _, page := range []string{"dashboard", "building-list", "building-detail"} {
		if got[page] != "app" {
			t.Errorf("layout of %q = %q, want \"app\"", page, got[page])
		}
	}
	if _, ok := got["login"]; ok {
		t.Error(`a block without data-layout must contribute nothing (delegates to defaultLayout)`)
	}

	if m := sitemapPageLayouts(nil); m != nil {
		t.Errorf("nil sitemap must map to nil, got %v", m)
	}
}
