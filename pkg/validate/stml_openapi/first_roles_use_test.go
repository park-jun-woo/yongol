//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestFirstRolesUse — 문서 순서 첫 data-roles 노드의 위치 경로/중첩 노드/미사용 시 "" 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestFirstRolesUse(t *testing.T) {
	sm := &stml.SitemapSpec{Navs: []stml.SitemapNav{
		{Items: []stml.SitemapNode{
			{Page: "dashboard", Label: "대시보드"}, // no roles
		}},
		{Items: []stml.SitemapNode{
			{Label: "관리", Children: []stml.SitemapNode{
				{Page: "billing", Label: "정산", Roles: []string{"admin"}}, // first data-roles in document order
			}},
			{Page: "audit", Label: "감사", Roles: []string{"admin", "manager"}}, // later use, ignored
		}},
	}}
	if got, want := firstRolesUse(sm), "nav[1] > 관리 > 정산"; got != want {
		t.Errorf("firstRolesUse = %q, want %q", got, want)
	}

	// no node carries data-roles → "" (TM-47 gate stays closed)
	none := &stml.SitemapSpec{Navs: []stml.SitemapNav{
		{Items: []stml.SitemapNode{{Page: "dashboard", Label: "대시보드"}}},
	}}
	if got := firstRolesUse(none); got != "" {
		t.Errorf("no roles: got %q, want \"\"", got)
	}

	// empty sitemap → ""
	if got := firstRolesUse(&stml.SitemapSpec{}); got != "" {
		t.Errorf("empty sitemap: got %q, want \"\"", got)
	}
}
