//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what TestAppendSitemapEntries — 라벨/page 폴백/(group) 표시명과 깊이별 위치 경로 누적 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestAppendSitemapEntries(t *testing.T) {
	nodes := []stml.SitemapNode{
		{Page: "dashboard", Label: "대시보드"},
		{Page: "member-list"}, // no label → page name as display
		{Label: "건물 관리", Children: []stml.SitemapNode{
			{Page: "building-detail", Label: "건물 상세"},
		}},
		{}, // neither label nor page → (group)
	}

	var entries []sitemapEntry
	appendSitemapEntries(nodes, "nav[0]", &entries)

	wantPaths := []string{
		"nav[0] > 대시보드",
		"nav[0] > member-list",
		"nav[0] > 건물 관리",
		"nav[0] > 건물 관리 > 건물 상세",
		"nav[0] > (group)",
	}
	if len(entries) != len(wantPaths) {
		t.Fatalf("expected %d entries, got %d: %+v", len(wantPaths), len(entries), entries)
	}
	for i, want := range wantPaths {
		if entries[i].Path != want {
			t.Errorf("entries[%d].Path = %q, want %q", i, entries[i].Path, want)
		}
	}
	if entries[3].Node.Page != "building-detail" {
		t.Errorf("entries[3].Node = %+v, want the nested child paired with its path", entries[3].Node)
	}
}
