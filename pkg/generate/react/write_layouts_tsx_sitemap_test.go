//ff:func feature=gen-react type=test control=sequence
//ff:what writeLayoutsTSX sitemap — 레이아웃별 메뉴 분리 ("" 블록 defaultLayout 위임 포함) 검증

package react

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestWriteLayoutsTSX_SitemapPerLayoutMenus(t *testing.T) {
	dir := t.TempDir()
	layouts := []stml.LayoutSpec{
		{Name: "app", File: "layouts/app.html", HasOutlet: true},
		{Name: "admin", File: "layouts/admin.html", HasOutlet: true},
	}
	patterns := map[string]string{"dashboard": "/dashboard", "admin-home": "/admin-home"}
	sitemap := &stml.SitemapSpec{Navs: []stml.SitemapNav{
		// no data-layout → delegates to defaultLayout "app"
		{Items: []stml.SitemapNode{{Page: "dashboard", Label: "대시보드", Menu: true}}},
		{Layout: "admin", Items: []stml.SitemapNode{{Page: "admin-home", Label: "관리 홈", Menu: true}}},
	}}

	if err := writeLayoutsTSX(dir, layouts, patterns, "", sitemap, "app", "", nil); err != nil {
		t.Fatal(err)
	}

	app, err := os.ReadFile(filepath.Join(dir, "layouts", "AppLayout.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	assertContains(t, string(app), `<NavLink to="/dashboard" end>대시보드</NavLink>`)
	assertNotContains(t, string(app), "관리 홈")

	admin, err := os.ReadFile(filepath.Join(dir, "layouts", "AdminLayout.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	assertContains(t, string(admin), `<NavLink to="/admin-home" end>관리 홈</NavLink>`)
	assertNotContains(t, string(admin), "대시보드")
}
