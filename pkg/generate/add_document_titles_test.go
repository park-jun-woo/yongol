//ff:func feature=generate type=test control=sequence
//ff:what addDocumentTitles — 라벨 폴백(페이지명) / 최초 등장 우선 / 그룹·외부 링크 무시 검증

package generate

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestAddDocumentTitles(t *testing.T) {
	titles := map[string]string{}
	addDocumentTitles([]stmlparser.SitemapNode{
		{Page: "building-list", Menu: true},                         // labelless — page-name fallback
		{Page: "building-list", Label: "두 번째 등장", Menu: true},       // duplicate (TM-40's ERROR) — first wins
		{Label: "외부", Href: "https://docs.example.com", Menu: true}, // external link — no entry
	}, "zenflow", titles)

	if got := titles["building-list"]; got != "building-list · zenflow" {
		t.Errorf("titles[building-list] = %q, want the page-name fallback with the first occurrence kept", got)
	}
	if len(titles) != 1 {
		t.Errorf("titles = %v, want a single entry", titles)
	}
}
