//ff:func feature=gen-react type=test control=sequence
//ff:what crumbFieldLayouts — 3단 배정 사슬별 레이아웃 귀속, 선언 없음/레이아웃 없음 nil·비귀속 검증

package react

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCrumbFieldLayouts(t *testing.T) {
	sitemap := &stml.SitemapSpec{Navs: []stml.SitemapNav{
		{Layout: "admin", Items: []stml.SitemapNode{
			{Page: "member-detail", CrumbField: "member_name"},
		}},
		{Items: []stml.SitemapNode{
			{Page: "building-detail", CrumbField: "building_name"},
			{Page: "report-detail", CrumbField: "report_title"},
			{Page: "home"},
		}},
	}}

	t.Run("three-step assignment chain", func(t *testing.T) {
		pages := []stml.PageSpec{
			{Name: "building-detail", Layout: "custom"}, // page data-layout wins
			{Name: "member-detail"},                     // sitemap nav data-layout
			{Name: "report-detail"},                     // defaultLayout
			{Name: "home"},                              // no crumb field — irrelevant
		}
		got := crumbFieldLayouts(pages, sitemap, "app")
		want := map[string]bool{"custom": true, "admin": true, "app": true}
		if len(got) != len(want) {
			t.Fatalf("crumbFieldLayouts = %v, want %v", got, want)
		}
		for k := range want {
			if !got[k] {
				t.Errorf("missing layout %q in %v", k, got)
			}
		}
	})

	t.Run("no declaration yields nil — byte-identity gate", func(t *testing.T) {
		plain := &stml.SitemapSpec{Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{{Page: "home"}}}}}
		if got := crumbFieldLayouts([]stml.PageSpec{{Name: "home"}}, plain, "app"); got != nil {
			t.Errorf("crumbFieldLayouts = %v, want nil", got)
		}
	})

	t.Run("layout-less page contributes nothing", func(t *testing.T) {
		got := crumbFieldLayouts([]stml.PageSpec{{Name: "report-detail"}}, sitemap, "")
		if got["report-detail"] || len(got) != 0 {
			t.Errorf("crumbFieldLayouts = %v, want empty for a bare page", got)
		}
	})

	t.Run("nil sitemap yields nil", func(t *testing.T) {
		if got := crumbFieldLayouts([]stml.PageSpec{{Name: "x"}}, nil, "app"); got != nil {
			t.Errorf("crumbFieldLayouts = %v, want nil", got)
		}
	})
}
