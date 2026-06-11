//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what sitemapDiags — 사이트맵 규칙 묶음 실행(기존 TM-39 + 동적 그룹 TM-48 동시 발화) 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestSitemapDiags(t *testing.T) {
	fs := makeFS(nil, nil)
	fs.Sitemap = &stml.SitemapSpec{FileName: "sitemap.html", Navs: []stml.SitemapNav{{Entry: true, Items: []stml.SitemapNode{
		{Page: "ghost-page", Label: "유령", Menu: true},
		{Label: "내 건물", Menu: true, Fetch: "ListMyBuildings", Each: "items", Link: "building-detail", LabelField: "name"},
	}}}}
	opMap := map[string]operationEntry{}
	diags := sitemapDiags(fs, opMap, nil)
	if got := countDiag(diags, "[TM-39]"); got != 1 {
		t.Errorf("expected the bundled TM-39 to fire once, got %d: %+v", got, diags)
	}
	if got := countDiag(diags, "[TM-48]"); got != 1 {
		t.Errorf("expected the bundled TM-48 to fire once, got %d: %+v", got, diags)
	}
}
