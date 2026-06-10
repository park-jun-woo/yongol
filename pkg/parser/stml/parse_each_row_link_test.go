//ff:func feature=stml-parse type=test control=sequence
//ff:what TestParseEachRowLink — 행 템플릿 요소의 data-link가 RowLink로 기록되고 셀 파싱은 유지됨을 검증

package stml

import (
	"strings"
	"testing"
)

func TestParseEachRowLink(t *testing.T) {
	input := `<main>
  <section data-fetch="ListBuildings">
    <ul data-each="buildings">
      <li data-link="building-detail" data-link-params="item.id -> BuildingID">
        <span data-bind="name"></span>
        <span data-bind="address"></span>
      </li>
    </ul>
  </section>
</main>`

	page, diags := ParseReader("building-list.html", strings.NewReader(input))
	if len(diags) > 0 {
		t.Fatal(diags)
	}
	eb := page.Fetches[0].Eaches[0]

	if eb.RowLink == nil {
		t.Fatal("RowLink not set from the item template's data-link")
	}
	if eb.RowLink.TargetPage != "building-detail" {
		t.Errorf("TargetPage = %q", eb.RowLink.TargetPage)
	}
	if len(eb.RowLink.Params) != 1 || eb.RowLink.Params[0].Segment != "BuildingID" {
		t.Errorf("Params = %+v", eb.RowLink.Params)
	}
	// The template's children keep parsing as row cells.
	if len(eb.Binds) != 2 || eb.Binds[0].Name != "name" || eb.Binds[1].Name != "address" {
		t.Errorf("Binds = %+v", eb.Binds)
	}
	// No RowLink without a template data-link.
	plain, pd := ParseReader("plain.html", strings.NewReader(
		`<main><section data-fetch="L"><ul data-each="xs"><li><span data-bind="a"></span></li></ul></section></main>`))
	if len(pd) > 0 {
		t.Fatal(pd)
	}
	if plain.Fetches[0].Eaches[0].RowLink != nil {
		t.Error("RowLink must be nil without a template data-link")
	}
}
