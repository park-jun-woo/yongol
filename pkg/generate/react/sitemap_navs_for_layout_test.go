//ff:func feature=gen-react type=test control=sequence
//ff:what sitemapNavsForLayout — data-layout 일치/"" defaultLayout 위임/타 레이아웃 분리/nil sitemap 검증

package react

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestSitemapNavsForLayout(t *testing.T) {
	sitemap := &stml.SitemapSpec{Navs: []stml.SitemapNav{
		{Layout: "app", Items: []stml.SitemapNode{{Page: "a", Menu: true}}},
		{Layout: "", Items: []stml.SitemapNode{{Page: "b", Menu: true}}},
		{Layout: "admin", Items: []stml.SitemapNode{{Page: "c", Menu: true}}},
		{Layout: "app", Items: []stml.SitemapNode{{Page: "d", Menu: true}}},
	}}

	t.Run("layout match keeps document order", func(t *testing.T) {
		navs := sitemapNavsForLayout(sitemap, "app", "admin")
		if len(navs) != 2 || navs[0].Items[0].Page != "a" || navs[1].Items[0].Page != "d" {
			t.Errorf("app navs = %+v", navs)
		}
	})

	t.Run("empty data-layout delegates to defaultLayout", func(t *testing.T) {
		navs := sitemapNavsForLayout(sitemap, "admin", "admin")
		if len(navs) != 2 || navs[0].Items[0].Page != "b" || navs[1].Items[0].Page != "c" {
			t.Errorf("admin navs = %+v", navs)
		}
	})

	t.Run("unassigned layout gets nothing", func(t *testing.T) {
		if navs := sitemapNavsForLayout(sitemap, "bare", "admin"); len(navs) != 0 {
			t.Errorf("bare navs = %+v", navs)
		}
	})

	t.Run("nil sitemap is nil", func(t *testing.T) {
		if navs := sitemapNavsForLayout(nil, "app", "app"); navs != nil {
			t.Errorf("expected nil, got %+v", navs)
		}
	})
}
