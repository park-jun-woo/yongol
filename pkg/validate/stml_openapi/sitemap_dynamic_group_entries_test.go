//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestSitemapDynamicGroupEntries — 동적 어휘 보유 노드만 위치 경로와 함께 평탄화되는지 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestSitemapDynamicGroupEntries(t *testing.T) {
	sm := &stml.SitemapSpec{FileName: "sitemap.html", Navs: []stml.SitemapNav{
		{Items: []stml.SitemapNode{
			{Page: "dashboard", Label: "대시보드"},
			{Label: "내 건물", Fetch: "ListMyBuildings", Each: "items", Link: "building-detail", LabelField: "building_name"},
		}},
		{Items: []stml.SitemapNode{
			{Label: "그룹", Children: []stml.SitemapNode{{Label: "중첩", Each: "rows"}}},
		}},
	}}
	entries := sitemapDynamicGroupEntries(sm)
	if len(entries) != 2 {
		t.Fatalf("entries = %+v, want the two vocab-carrying nodes", entries)
	}
	if entries[0].Node.Fetch != "ListMyBuildings" || entries[0].Path != "nav[0] > 내 건물" {
		t.Errorf("entry[0] = %+v", entries[0])
	}
	if entries[1].Node.Each != "rows" || entries[1].Path != "nav[1] > 그룹 > 중첩" {
		t.Errorf("entry[1] = %+v", entries[1])
	}
}
