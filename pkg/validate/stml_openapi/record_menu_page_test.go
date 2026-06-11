//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what recordMenuPage — 그룹 스킵 / 렌더→루트 / 차단 사유 최초 기록 보존 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRecordMenuPage(t *testing.T) {
	g := &pageGraph{Roots: map[string]bool{}, InSitemap: map[string]bool{}, MenuBlocked: map[string]string{}}

	recordMenuPage(stml.SitemapNode{Label: "그룹", Menu: true}, "", g)
	if len(g.InSitemap) != 0 {
		t.Errorf("a group without data-page must record nothing, got %+v", g.InSitemap)
	}

	recordMenuPage(stml.SitemapNode{Page: "dashboard", Menu: true}, "", g)
	if !g.InSitemap["dashboard"] || !g.Roots["dashboard"] {
		t.Error("a rendered page should be listed and a root")
	}

	recordMenuPage(stml.SitemapNode{Page: "member-list", Menu: true}, "first reason", g)
	recordMenuPage(stml.SitemapNode{Page: "member-list", Menu: true}, "second reason", g)
	if g.Roots["member-list"] {
		t.Error("a blocked page must not be a root")
	}
	if g.MenuBlocked["member-list"] != "first reason" {
		t.Errorf("MenuBlocked = %q, want the first recorded reason kept", g.MenuBlocked["member-list"])
	}
}
