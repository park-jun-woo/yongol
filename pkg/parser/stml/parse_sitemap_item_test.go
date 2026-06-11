//ff:func feature=stml-parse type=test control=sequence
//ff:what TestParseSitemapItem — li 속성(data-page/index/menu/icon/crumb-field)과 라벨/자식 ul/외부 링크 변환 검증

package stml

import "testing"

func TestParseSitemapItem(t *testing.T) {
	t.Run("full attribute set", func(t *testing.T) {
		li := firstElementNode(t, `<li data-page="dashboard" data-index data-icon="home">대시보드</li>`, "li")
		var spec SitemapSpec
		node := parseSitemapItem(li, &spec)
		if node.Page != "dashboard" || node.Label != "대시보드" || !node.Index || node.Icon != "home" {
			t.Errorf("node = %+v", node)
		}
		if !node.Menu {
			t.Errorf("Menu should default to true")
		}
		if node.CrumbField != "" {
			t.Errorf("CrumbField = %q, want empty without data-crumb-field", node.CrumbField)
		}
	})

	t.Run("data-crumb-field is parsed first-class (Phase006)", func(t *testing.T) {
		li := firstElementNode(t, `<li data-page="building-detail" data-crumb-field="building_name">건물 상세</li>`, "li")
		var spec SitemapSpec
		node := parseSitemapItem(li, &spec)
		if node.CrumbField != "building_name" {
			t.Errorf("CrumbField = %q, want building_name", node.CrumbField)
		}
	})

	t.Run("dynamic group vocabulary is parsed first-class (Phase007)", func(t *testing.T) {
		li := firstElementNode(t, `<li>내 건물<ul data-fetch="ListMyBuildings" data-each="items" data-link="building-detail" data-link-params="item.building_id -> BuildingID" data-label-field="building_name"></ul></li>`, "li")
		var spec SitemapSpec
		node := parseSitemapItem(li, &spec)
		if node.Fetch != "ListMyBuildings" || node.Each != "items" || node.Link != "building-detail" || node.LabelField != "building_name" {
			t.Errorf("dynamic node = %+v, want the five graduated fields", node)
		}
		if node.LinkParamsRaw != "item.building_id -> BuildingID" {
			t.Errorf("LinkParamsRaw = %q", node.LinkParamsRaw)
		}
		if len(node.LinkParams) != 1 || node.LinkParams[0].Source != "item.building_id" || node.LinkParams[0].Segment != "BuildingID" {
			t.Errorf("LinkParams = %+v", node.LinkParams)
		}
	})

	t.Run("data-menu false hides the menu entry", func(t *testing.T) {
		li := firstElementNode(t, `<li data-page="member-list" data-menu="false">멤버</li>`, "li")
		var spec SitemapSpec
		node := parseSitemapItem(li, &spec)
		if node.Menu {
			t.Errorf("Menu = true, want false for data-menu=\"false\"")
		}
		if node.Index {
			t.Errorf("Index = true, want false without data-index")
		}
	})

	t.Run("group label with nested ul children", func(t *testing.T) {
		li := firstElementNode(t, `<li>건물 관리<ul><li data-page="building-list">건물 목록</li></ul></li>`, "li")
		var spec SitemapSpec
		node := parseSitemapItem(li, &spec)
		if node.Page != "" || node.Label != "건물 관리" {
			t.Errorf("group node = %+v, want page-less label", node)
		}
		if len(node.Children) != 1 || node.Children[0].Page != "building-list" {
			t.Errorf("Children = %+v", node.Children)
		}
	})

	t.Run("external link child fills href", func(t *testing.T) {
		li := firstElementNode(t, `<li><a href="https://docs.example.com">매뉴얼</a></li>`, "li")
		var spec SitemapSpec
		node := parseSitemapItem(li, &spec)
		if node.Href != "https://docs.example.com" || node.Label != "매뉴얼" {
			t.Errorf("node = %+v", node)
		}
	})
}
