//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what collectLinkRefs — 컨텍스트(each 내부 여부·item 스키마) 부착 수집 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCollectLinkRefs(t *testing.T) {
	page, parseDiags := stml.ParseReader("page.html", strings.NewReader(`<main>
  <section data-fetch="ListBuildings">
    <ul data-each="buildings">
      <li data-link="building-detail" data-link-params="item.id -> BuildingID">
        <span data-bind="name"></span>
      </li>
    </ul>
  </section>
  <a data-link="settings">설정</a>
</main>`))
	if len(parseDiags) > 0 {
		t.Fatalf("unexpected parse diags: %v", parseDiags)
	}
	raif := map[string]map[string]map[string]bool{
		"ListBuildings": {"buildings": {"id": true, "name": true}},
	}
	var refs []linkRefCtx
	collectLinkRefs(page.Children, "", nil, false, raif, &refs)
	if len(refs) != 2 {
		t.Fatalf("got %d refs, want 2: %+v", len(refs), refs)
	}
	byTarget := map[string]linkRefCtx{}
	for _, r := range refs {
		byTarget[r.Link.TargetPage] = r
	}
	row, ok := byTarget["building-detail"]
	if !ok || !row.InEach {
		t.Errorf("row link: ok=%v InEach=%v", ok, row.InEach)
	}
	if row.ItemFields == nil || !row.ItemFields["id"] {
		t.Errorf("row link item fields = %+v", row.ItemFields)
	}
	nav, ok := byTarget["settings"]
	if !ok || nav.InEach || nav.ItemFields != nil {
		t.Errorf("nav link: ok=%v InEach=%v fields=%+v", ok, nav.InEach, nav.ItemFields)
	}
}
