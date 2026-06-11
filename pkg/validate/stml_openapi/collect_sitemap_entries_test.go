//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what TestCollectSitemapEntries — 복수 nav 블록을 nav[i] 접두사로 문서 순서 평탄화 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCollectSitemapEntries(t *testing.T) {
	sm := &stml.SitemapSpec{
		FileName: "sitemap.html",
		Navs: []stml.SitemapNav{
			{Items: []stml.SitemapNode{
				{Page: "dashboard", Label: "대시보드", Children: []stml.SitemapNode{
					{Page: "building-list", Label: "건물 목록"},
				}},
			}},
			{Items: []stml.SitemapNode{
				{Page: "login", Label: "로그인"},
			}},
		},
	}

	entries := collectSitemapEntries(sm)
	wantPaths := []string{
		"nav[0] > 대시보드",
		"nav[0] > 대시보드 > 건물 목록",
		"nav[1] > 로그인",
	}
	if len(entries) != len(wantPaths) {
		t.Fatalf("expected %d entries, got %d: %+v", len(wantPaths), len(entries), entries)
	}
	for i, want := range wantPaths {
		if entries[i].Path != want {
			t.Errorf("entries[%d].Path = %q, want %q", i, entries[i].Path, want)
		}
	}
}
