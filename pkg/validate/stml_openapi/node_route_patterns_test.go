//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what nodeRoutePatterns — 페이지 노드 해석 / 그룹·외부 링크 nil / 미실존 페이지 nil 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestNodeRoutePatterns(t *testing.T) {
	pages := []stml.PageSpec{
		{Name: "building-detail", FileName: "building-detail.html", Route: "/buildings/:BuildingID"},
	}

	t.Run("page node resolves to RoutePaths", func(t *testing.T) {
		got := nodeRoutePatterns(stml.SitemapNode{Page: "building-detail", Menu: true}, pages)
		if len(got) != 1 || got[0] != "/buildings/:BuildingID" {
			t.Errorf("patterns = %v, want [/buildings/:BuildingID]", got)
		}
	})

	t.Run("group node is nil", func(t *testing.T) {
		if got := nodeRoutePatterns(stml.SitemapNode{Label: "관리", Menu: true}, pages); got != nil {
			t.Errorf("patterns = %v, want nil for a group", got)
		}
	})

	t.Run("nonexistent page is nil", func(t *testing.T) {
		if got := nodeRoutePatterns(stml.SitemapNode{Page: "ghost", Menu: true}, pages); got != nil {
			t.Errorf("patterns = %v, want nil for a ghost page", got)
		}
	})
}
