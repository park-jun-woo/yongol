//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what Run 경유 TM-46 — sitemap data-roles 오타 role ERROR 발화, sitemap 부재 시 침묵 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM46_RunFires(t *testing.T) {
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/members": getOp("ListMembers", nil, nil),
	})
	pages := []stml.PageSpec{
		{Name: "member-list", FileName: "member-list.html", Fetches: []stml.FetchBlock{{OperationID: "ListMembers"}}},
	}
	fs := makeAuthFS(pages, doc, "cookie")
	fs.Manifest.Backend.Auth.Roles = []string{"member", "admin"}
	fs.Sitemap = &stml.SitemapSpec{
		FileName: "sitemap.html",
		Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{
			{Page: "member-list", Label: "멤버", Index: true, Menu: true, Roles: []string{"adimn"}},
		}}},
	}

	if got := countDiag(Run(fs), "[TM-46]"); got != 1 {
		t.Errorf("expected 1 TM-46 via Run, got %d: %+v", got, Run(fs))
	}

	fs.Sitemap = nil
	if got := countDiag(Run(fs), "[TM-46]"); got != 0 {
		t.Errorf("expected 0 TM-46 via Run without a sitemap, got %d", got)
	}
}
