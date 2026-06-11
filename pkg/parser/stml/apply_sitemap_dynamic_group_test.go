//ff:func feature=stml-parse type=test control=sequence
//ff:what TestApplySitemapDynamicGroup — 동적 어휘 없는 ul 무시·5종 승격·첫 동적 ul 승자·link-params 파싱 검증

package stml

import "testing"

func TestApplySitemapDynamicGroup(t *testing.T) {
	t.Run("a plain ul contributes nothing", func(t *testing.T) {
		ul := firstElementNode(t, `<ul><li data-page="a">A</li></ul>`, "ul")
		var node SitemapNode
		applySitemapDynamicGroup(ul, &node)
		if node.Fetch != "" || node.Each != "" || node.Link != "" || node.LinkParamsRaw != "" || node.LabelField != "" {
			t.Errorf("node = %+v, want no dynamic fields", node)
		}
	})

	t.Run("all five attributes graduate onto the node", func(t *testing.T) {
		ul := firstElementNode(t, `<ul data-fetch="ListMyBuildings" data-each="items" data-link="building-detail" data-link-params="item.building_id -> BuildingID" data-label-field="building_name"></ul>`, "ul")
		var node SitemapNode
		applySitemapDynamicGroup(ul, &node)
		if node.Fetch != "ListMyBuildings" || node.Each != "items" || node.Link != "building-detail" || node.LabelField != "building_name" {
			t.Errorf("node = %+v", node)
		}
		if len(node.LinkParams) != 1 || node.LinkParams[0].Source != "item.building_id" || node.LinkParams[0].Segment != "BuildingID" {
			t.Errorf("LinkParams = %+v", node.LinkParams)
		}
	})

	t.Run("the first dynamic ul wins", func(t *testing.T) {
		first := firstElementNode(t, `<ul data-fetch="First" data-each="items"></ul>`, "ul")
		second := firstElementNode(t, `<ul data-fetch="Second" data-each="rows"></ul>`, "ul")
		var node SitemapNode
		applySitemapDynamicGroup(first, &node)
		applySitemapDynamicGroup(second, &node)
		if node.Fetch != "First" || node.Each != "items" {
			t.Errorf("node = %+v, want the first dynamic ul preserved", node)
		}
	})

	t.Run("an invalid data-link-params keeps the raw value without bindings", func(t *testing.T) {
		ul := firstElementNode(t, `<ul data-fetch="ListThings" data-link-params="nonsense"></ul>`, "ul")
		var node SitemapNode
		applySitemapDynamicGroup(ul, &node)
		if node.LinkParamsRaw != "nonsense" || len(node.LinkParams) != 0 {
			t.Errorf("node = raw %q params %+v", node.LinkParamsRaw, node.LinkParams)
		}
	})
}
