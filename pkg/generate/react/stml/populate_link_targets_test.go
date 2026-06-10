//ff:func feature=stml-gen type=test control=sequence
//ff:what TestPopulateLinkTargets — 페이지 내 모든 LinkRef(행 링크 포함)에 대상 패턴 설정 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestPopulateLinkTargets(t *testing.T) {
	page, parseDiags := stmlparser.ParseReader("building-list.html", strings.NewReader(`<main>
  <section data-fetch="ListBuildings">
    <ul data-each="buildings">
      <li data-link="building-detail" data-link-params="item.id -> BuildingID">
        <span data-bind="name"></span>
      </li>
    </ul>
  </section>
  <nav>
    <a data-link="settings">설정</a>
  </nav>
</main>`))
	if len(parseDiags) > 0 {
		t.Fatal(parseDiags)
	}

	patterns := map[string]string{
		"building-detail": "/buildings/:BuildingID",
		"settings":        "/settings",
	}
	populateLinkTargets(&page, patterns)

	var rowPattern, navPattern string
	var walk func(children []stmlparser.ChildNode)
	walk = func(children []stmlparser.ChildNode) {
		for _, ch := range children {
			switch ch.Kind {
			case "link":
				navPattern = ch.Link.TargetPattern
			case "each":
				if ch.Each.RowLink != nil {
					rowPattern = ch.Each.RowLink.TargetPattern
				}
			case "fetch":
				walk(ch.Fetch.Children)
			case "static":
				walk(ch.Static.Children)
			}
		}
	}
	walk(page.Children)

	if rowPattern != "/buildings/:BuildingID" {
		t.Errorf("RowLink.TargetPattern = %q", rowPattern)
	}
	if navPattern != "/settings" {
		t.Errorf("nav link TargetPattern = %q", navPattern)
	}

	// nil map is a no-op.
	populateLinkTargets(&page, nil)
}
