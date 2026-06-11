//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TM-13 — sitemap nav 블록 data-layout 으로만 배정된 레이아웃은 used (오탐 차단) 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM13_SitemapLayoutCountsAsUsed(t *testing.T) {
	pages := []stml.PageSpec{{Name: "admin-home", FileName: "admin-home.html"}}
	layouts := []stml.LayoutSpec{{Name: "admin", File: "layouts/admin.html"}}
	sitemap := &stml.SitemapSpec{
		FileName: "sitemap.html",
		Navs: []stml.SitemapNav{{Layout: "admin", Items: []stml.SitemapNode{
			{Page: "admin-home", Label: "관리 홈", Menu: true},
		}}},
	}

	// The Phase003 emitter assigns admin-home to "admin" via the sitemap
	// block, so flagging the layout unused would be a validation/emission
	// drift — it must count as used.
	if diags := tm13UnusedLayout(pages, layouts, "", sitemap); len(diags) != 0 {
		t.Errorf("sitemap-assigned layout must be used, got %+v", diags)
	}

	// Without the sitemap the same fixture is genuinely unused.
	if diags := tm13UnusedLayout(pages, layouts, "", nil); countDiag(diags, "[TM-13]") != 1 {
		t.Errorf("expected 1 TM-13 without the sitemap, got %+v", diags)
	}
}
