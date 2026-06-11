//ff:func feature=stml-parse type=test control=sequence
//ff:what TestParseSitemapReader_DynamicGroup — 동적 그룹 어휘 5종의 1급 파싱(첫 ul 승자·잘못된 link-params 보존)과 roles/crumb-field 승계 검증

package stml

import (
	"strings"
	"testing"
)

func TestParseSitemapReader_DynamicGroup(t *testing.T) {
	src := `
<nav data-sitemap>
  <ul>
    <li data-page="admin-home" data-roles="admin">관리</li>
    <li data-page="building-detail" data-crumb-field="building_name">건물 상세</li>
    <li>내 건물
      <ul data-fetch="ListMyBuildings" data-each="items" data-link="building-detail" data-link-params="item.building_id -> BuildingID" data-label-field="building_name"></ul>
    </li>
    <li>고장난 그룹
      <ul data-fetch="ListThings" data-each="items" data-link="thing-detail" data-link-params="broken" data-label-field="name"></ul>
    </li>
  </ul>
</nav>`
	spec, diags := ParseSitemapReader("sitemap.html", strings.NewReader(src))
	if len(diags) != 0 {
		t.Fatalf("expected no diags, got %+v", diags)
	}
	items := spec.Navs[0].Items
	// data-roles (Phase005) and data-crumb-field (Phase006) stay first-class.
	if got := items[0].Roles; len(got) != 1 || got[0] != "admin" {
		t.Errorf("admin-home Roles = %v, want [admin]", got)
	}
	if got := items[1].CrumbField; got != "building_name" {
		t.Errorf("building-detail CrumbField = %q, want building_name", got)
	}
	// Dynamic group vocabulary (Phase007) graduates onto the group node.
	group := items[2]
	if group.Fetch != "ListMyBuildings" || group.Each != "items" || group.Link != "building-detail" || group.LabelField != "building_name" {
		t.Errorf("dynamic group = %+v", group)
	}
	if len(group.LinkParams) != 1 || group.LinkParams[0].Source != "item.building_id" || group.LinkParams[0].Segment != "BuildingID" {
		t.Errorf("LinkParams = %+v", group.LinkParams)
	}
	// A syntactically invalid data-link-params keeps the raw value (TM-32
	// re-parses it for the diagnostic) and parses to no bindings.
	broken := items[3]
	if broken.LinkParamsRaw != "broken" || len(broken.LinkParams) != 0 {
		t.Errorf("broken group = raw %q params %+v, want raw preserved and no bindings", broken.LinkParamsRaw, broken.LinkParams)
	}
}
