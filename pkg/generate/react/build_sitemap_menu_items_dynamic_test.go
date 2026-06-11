//ff:func feature=gen-react type=test control=sequence
//ff:what TestBuildSitemapMenuItemsDynamic — 완전 동적 그룹의 fetch 배선·LinkToAttr 경로 치환·키 선택, 불완전 그룹의 정적 유지 검증

package react

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestBuildSitemapMenuItemsDynamic(t *testing.T) {
	routePatterns := map[string]string{"building-detail": "/buildings/:BuildingID"}
	node := stml.SitemapNode{
		Label: "내 건물", Menu: true,
		Fetch: "ListMyBuildings", Each: "items", Link: "building-detail",
		LinkParamsRaw: "item.building_id -> BuildingID",
		LinkParams:    []stml.LinkParamBind{{Source: "item.building_id", Segment: "BuildingID"}},
		LabelField:    "building_name",
	}

	t.Run("a complete group carries the fetch wiring and the substituted path", func(t *testing.T) {
		items := buildSitemapMenuItems([]stml.SitemapNode{node}, 1, routePatterns)
		if len(items) != 1 {
			t.Fatalf("items = %+v", items)
		}
		it := items[0]
		if it.Kind != "group" || it.Fetch != "ListMyBuildings" || it.Each != "items" || it.LabelField != "building_name" {
			t.Errorf("item = %+v", it)
		}
		if it.ItemToAttr != "to={`/buildings/${item.building_id}`}" {
			t.Errorf("ItemToAttr = %q, want the page data-link substitution", it.ItemToAttr)
		}
		if it.ItemKey != "item.building_id" {
			t.Errorf("ItemKey = %q", it.ItemKey)
		}
	})

	t.Run("an incomplete group stays static", func(t *testing.T) {
		incomplete := node
		incomplete.LabelField = ""
		items := buildSitemapMenuItems([]stml.SitemapNode{incomplete}, 1, routePatterns)
		if len(items) != 1 || items[0].Fetch != "" || items[0].ItemToAttr != "" {
			t.Errorf("items = %+v, want a plain static group", items)
		}
	})
}
