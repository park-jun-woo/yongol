//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what Run 경유 TM-51 — sitemap 메뉴 렌더 + 호스트 레이아웃 전무 WARNING 발화, 레이아웃 존재·sitemap 부재 시 침묵 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM51_RunFires(t *testing.T) {
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/home": getOp("GetHome", nil, nil),
	})
	pages := []stml.PageSpec{{Name: "home", FileName: "home.html"}}
	fs := makeFS(pages, doc)
	fs.Sitemap = &stml.SitemapSpec{
		FileName: "sitemap.html",
		Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{
			{Page: "home", Label: "Home", Index: true, Menu: true},
		}}},
	}

	if got := countDiag(Run(fs), "[TM-51]"); got != 1 {
		t.Errorf("expected 1 TM-51 via Run, got %d: %+v", got, Run(fs))
	}

	// A layout host silences it (and engages the layout rules instead).
	fs.Layouts = []stml.LayoutSpec{{Name: "app", HasOutlet: true}}
	if got := countDiag(Run(fs), "[TM-51]"); got != 0 {
		t.Errorf("expected 0 TM-51 via Run with a layout, got %d", got)
	}
	fs.Layouts = nil

	// No double diagnosis: defaultLayout set + empty layouts/ already
	// raises TM-12 ERROR, so TM-51 stays silent.
	fs.Manifest.Frontend.DefaultLayout = "app"
	if got := countDiag(Run(fs), "[TM-12]"); got != 1 {
		t.Errorf("expected 1 TM-12 via Run with a missing defaultLayout, got %d", got)
	}
	if got := countDiag(Run(fs), "[TM-51]"); got != 0 {
		t.Errorf("expected 0 TM-51 via Run while TM-12 owns it, got %d", got)
	}
	fs.Manifest.Frontend.DefaultLayout = ""

	// No double diagnosis: a nav data-layout + empty layouts/ already
	// raises TM-41 ERROR, so TM-51 stays silent.
	fs.Sitemap.Navs[0].Layout = "app"
	if got := countDiag(Run(fs), "[TM-41]"); got != 1 {
		t.Errorf("expected 1 TM-41 via Run with a missing nav data-layout, got %d", got)
	}
	if got := countDiag(Run(fs), "[TM-51]"); got != 0 {
		t.Errorf("expected 0 TM-51 via Run while TM-41 owns it, got %d", got)
	}
	fs.Sitemap.Navs[0].Layout = ""

	// Without a sitemap, sitemapDiags never runs — TM-51 stays silent.
	fs.Sitemap = nil
	if got := countDiag(Run(fs), "[TM-51]"); got != 0 {
		t.Errorf("expected 0 TM-51 via Run without a sitemap, got %d", got)
	}
}
