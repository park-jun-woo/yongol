//ff:func feature=gen-react type=test control=sequence
//ff:what breadcrumbCrumbFor — MenuRenderable 페이지 href / 그룹·외부 링크·필수 파라미터 라벨만 / 라벨 폴백 검증

package react

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestBreadcrumbCrumbFor(t *testing.T) {
	routePatterns := map[string]string{
		"building-list":   "/buildings",
		"building-detail": "/buildings/:BuildingID",
	}

	t.Run("menu-renderable page links", func(t *testing.T) {
		c := breadcrumbCrumbFor(stml.SitemapNode{Page: "building-list", Label: "건물 목록", Menu: true}, 2, routePatterns)
		if c.Label != "건물 목록" || c.Href != "/buildings" {
			t.Errorf("crumb = %+v, want linked 건물 목록", c)
		}
	})

	t.Run("required-param page stays label-only", func(t *testing.T) {
		c := breadcrumbCrumbFor(stml.SitemapNode{Page: "building-detail", Label: "건물 상세", Menu: true}, 2, routePatterns)
		if c.Label != "건물 상세" || c.Href != "" {
			t.Errorf("crumb = %+v, want label-only 건물 상세", c)
		}
	})

	t.Run("group label has no href", func(t *testing.T) {
		c := breadcrumbCrumbFor(stml.SitemapNode{Label: "건물 관리", Menu: true}, 1, routePatterns)
		if c.Label != "건물 관리" || c.Href != "" {
			t.Errorf("crumb = %+v, want label-only group", c)
		}
	})

	t.Run("external link has no href", func(t *testing.T) {
		c := breadcrumbCrumbFor(stml.SitemapNode{Label: "매뉴얼", Href: "https://docs.example.com", Menu: true}, 1, routePatterns)
		if c.Label != "매뉴얼" || c.Href != "" {
			t.Errorf("crumb = %+v, want label-only external entry", c)
		}
	})

	t.Run("labelless page falls back to the page name", func(t *testing.T) {
		c := breadcrumbCrumbFor(stml.SitemapNode{Page: "building-list", Menu: true}, 2, routePatterns)
		if c.Label != "building-list" {
			t.Errorf("Label = %q, want the page-name fallback", c.Label)
		}
	})
}
