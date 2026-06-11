//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what TestCollectConsumedOps — 전 페이지 data-fetch/data-action + 사이트맵 동적 그룹 data-fetch operationId 집합 수집
package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCollectConsumedOps(t *testing.T) {
	// Empty input.
	if got := collectConsumedOps(nil, nil, nil, "", nil); len(got) != 0 {
		t.Fatalf("expected empty, got %+v", got)
	}

	pages := []stml.PageSpec{{
		Name:    "page",
		Fetches: []stml.FetchBlock{{OperationID: "ListItems"}},
		Actions: []stml.ActionBlock{{OperationID: "CreateItem"}},
	}}
	out := collectConsumedOps(pages, nil, nil, "", nil)
	for _, id := range []string{"ListItems", "CreateItem"} {
		if _, ok := out[id]; !ok {
			t.Errorf("missing operationId %q", id)
		}
	}

	// A sitemap dynamic menu group's data-fetch is a real consumer too
	// (plans/stml/sitemap Phase007 — the layout emits a useQuery for it).
	sm := &stml.SitemapSpec{FileName: "sitemap.html", Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{
		{Label: "내 건물", Fetch: "ListMyBuildings", Each: "items", Link: "building-detail", LabelField: "name"},
	}}}}
	out = collectConsumedOps(pages, nil, sm, "", nil)
	if _, ok := out["ListMyBuildings"]; !ok {
		t.Errorf("missing sitemap dynamic-group operationId ListMyBuildings: %+v", out)
	}
}
