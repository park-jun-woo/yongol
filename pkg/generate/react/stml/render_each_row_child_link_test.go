//ff:func feature=stml-gen type=test control=sequence
//ff:what 행 자식 data-link — 행 액션과 같은 후행 셀(<TD><Link>)로 방출됨을 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderEachRowChildLink(t *testing.T) {
	page, parseDiags := stmlparser.ParseReader("building-list.html", strings.NewReader(`<main>
  <section data-fetch="ListBuildings">
    <ul data-each="buildings">
      <li>
        <span data-bind="name"></span>
        <a data-link="building-detail" data-link-params="item.id -> BuildingID">상세</a>
      </li>
    </ul>
  </section>
</main>`))
	if len(parseDiags) > 0 {
		t.Fatal(parseDiags)
	}
	opt := GenerateOptions{
		ResponseArrayItemFields: map[string]map[string]map[string]bool{
			"ListBuildings": {"buildings": {"id": true, "name": true}},
		},
		RoutePatterns: map[string]string{
			"building-detail": "/buildings/:BuildingID",
		},
	}
	code := GeneratePage(page, "", opt)

	// The row-child link gets its own trailing cell with an empty header,
	// and the field cell stays unwrapped (no RowLink).
	assertContains(t, code, "<TH></TH>")
	assertContains(t, code, "<Link to={`/buildings/${item.id}`}>상세</Link>")
	assertContains(t, code, "<TD>{item.name}</TD>")
}
