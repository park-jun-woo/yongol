//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what addEntryRoots — entry 블록 전 깊이 페이지 루트 편입 + 미실존 페이지 제외 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestAddEntryRoots(t *testing.T) {
	g := &pageGraph{Roots: map[string]bool{}}
	names := map[string]bool{"login": true, "signup": true}
	items := []stml.SitemapNode{
		{Page: "login", Menu: true, Children: []stml.SitemapNode{
			{Page: "signup", Menu: true},
		}},
		{Page: "ghost", Menu: true},
	}
	addEntryRoots(items, names, g)
	if !g.Roots["login"] || !g.Roots["signup"] {
		t.Errorf("Roots = %+v, want login and signup (all depths)", g.Roots)
	}
	if g.Roots["ghost"] {
		t.Error("a nonexistent page must not become a root")
	}
}
