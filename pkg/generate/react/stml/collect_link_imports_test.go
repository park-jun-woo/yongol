//ff:func feature=stml-gen type=test control=sequence
//ff:what TestCollectLinkImports — 링크 노드의 Link/useParams 임포트 플래그 수집 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCollectLinkImports(t *testing.T) {
	// Row link with an item.* source → Link only.
	page, parseDiags := stmlparser.ParseReader("building-list.html", strings.NewReader(`<main>
  <section data-fetch="ListBuildings">
    <ul data-each="buildings">
      <li data-link="building-detail" data-link-params="item.id -> BuildingID">
        <span data-bind="name"></span>
      </li>
    </ul>
  </section>
</main>`))
	if len(parseDiags) > 0 {
		t.Fatal(parseDiags)
	}
	var is importSet
	collectLinkImports(page.Children, &is)
	if !is.useLink {
		t.Error("useLink not set for a row link")
	}
	if is.useParams {
		t.Error("useParams must not be set for item.* sources")
	}

	// Static link with a route.* source → Link + useParams.
	page2, parseDiags2 := stmlparser.ParseReader("building-detail.html", strings.NewReader(`<main>
  <a data-link="unit-list" data-link-params="route.BuildingID -> BuildingID">세대 목록</a>
</main>`))
	if len(parseDiags2) > 0 {
		t.Fatal(parseDiags2)
	}
	var is2 importSet
	collectLinkImports(page2.Children, &is2)
	if !is2.useLink || !is2.useParams {
		t.Errorf("useLink=%v useParams=%v, want both true", is2.useLink, is2.useParams)
	}

	// No links → no flags.
	var is3 importSet
	collectLinkImports(nil, &is3)
	if is3.useLink || is3.useParams {
		t.Errorf("empty: got %+v", is3)
	}
}
