//ff:func feature=stml-gen type=test control=sequence
//ff:what 목록→상세 행 링크 페이지 생성 — Link 임포트·행 셀 Link 래핑·정적 링크 방출 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestGenerateLinkPage(t *testing.T) {
	page, parseDiags := stmlparser.ParseReader("building-list.html", strings.NewReader(`<main>
  <section data-fetch="ListBuildings">
    <ul data-each="buildings">
      <li data-link="building-detail" data-link-params="item.id -> BuildingID">
        <span data-bind="name"></span>
        <span data-bind="address"></span>
      </li>
    </ul>
  </section>
  <nav>
    <a data-link="settings-parsing-rules">파싱 규칙</a>
  </nav>
</main>`))
	if len(parseDiags) > 0 {
		t.Fatal(parseDiags)
	}
	opt := GenerateOptions{
		ResponseArrayItemFields: map[string]map[string]map[string]bool{
			"ListBuildings": {"buildings": {"id": true, "name": true, "address": true}},
		},
		RoutePatterns: map[string]string{
			"building-list":          "/building-list",
			"building-detail":        "/buildings/:BuildingID/:PhotoID?",
			"settings-parsing-rules": "/settings-parsing-rules",
		},
	}
	code := GeneratePage(page, "", opt)

	// react-router-dom Link import is emitted.
	assertContains(t, code, "import { Link } from 'react-router-dom'")
	// Whole-row link: every field cell's content is wrapped — the emitted
	// path is the target page's resolved route, not the page name, and the
	// unmapped optional :PhotoID? segment is omitted.
	assertContains(t, code, "<TD><Link to={`/buildings/${item.id}`}>{item.name}</Link></TD>")
	assertContains(t, code, "<TD><Link to={`/buildings/${item.id}`}>{item.address}</Link></TD>")
	// Static navigation link renders as a plain-path Link.
	assertContains(t, code, `<Link to="/settings-parsing-rules">파싱 규칙</Link>`)
	// No row-link page state leaks into pages without links.
	plain, pd := stmlparser.ParseReader("plain-list.html", strings.NewReader(`<main>
  <section data-fetch="ListBuildings">
    <ul data-each="buildings">
      <li><span data-bind="name"></span></li>
    </ul>
  </section>
</main>`))
	if len(pd) > 0 {
		t.Fatal(pd)
	}
	plainCode := GeneratePage(plain, "", opt)
	assertNotContains(t, plainCode, "react-router-dom")
	assertNotContains(t, plainCode, "<Link")
}
