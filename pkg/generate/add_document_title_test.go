//ff:func feature=generate type=test control=sequence
//ff:what addDocumentTitle — 라벨·앱명 결합 / 라벨 폴백 / 기존 항목 보존 / 그룹 무시 검증

package generate

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestAddDocumentTitle(t *testing.T) {
	titles := map[string]string{"dashboard": "대시보드 · zenflow"}

	addDocumentTitle(stmlparser.SitemapNode{Page: "building-list", Label: "건물 목록"}, "zenflow", titles)
	if titles["building-list"] != "건물 목록 · zenflow" {
		t.Errorf("titles[building-list] = %q", titles["building-list"])
	}

	addDocumentTitle(stmlparser.SitemapNode{Page: "dashboard", Label: "다른 라벨"}, "zenflow", titles)
	if titles["dashboard"] != "대시보드 · zenflow" {
		t.Errorf("first occurrence must win, got %q", titles["dashboard"])
	}

	addDocumentTitle(stmlparser.SitemapNode{Page: "settings"}, "", titles)
	if titles["settings"] != "settings" {
		t.Errorf("want bare page-name fallback without an app name, got %q", titles["settings"])
	}

	addDocumentTitle(stmlparser.SitemapNode{Label: "그룹"}, "zenflow", titles)
	if len(titles) != 3 {
		t.Errorf("group label must record nothing, got %v", titles)
	}
}
