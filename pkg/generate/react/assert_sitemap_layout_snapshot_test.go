//ff:func feature=gen-react type=test-helper control=sequence
//ff:what assertSitemapLayoutSnapshot — sitemap 메뉴 레이아웃 TSX 전체 스냅샷 비교 헬퍼

package react

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// assertSitemapLayoutSnapshot compares the full layout TSX emitted for a
// sitemap-derived menu against the pinned snapshot: NavLink/useLocation/
// lucide/Breadcrumb imports, the pathname const, the 2-level menu with
// the ancestor-highlight className, and the Breadcrumb above the Outlet
// (plans/stml/sitemap Phase004).
func assertSitemapLayoutSnapshot(t *testing.T, layout stml.LayoutSpec, menu *sitemapMenu) {
	t.Helper()
	got := renderLayoutTSX("AppLayout", layout, nil, "", menu)
	want := "import { NavLink, Outlet, useLocation } from 'react-router-dom'\n" +
		"import { LayoutDashboard } from 'lucide-react'\n" +
		"import { Breadcrumb } from '@/components/ui/Breadcrumb'\n" +
		"\n" +
		"export default function AppLayout() {\n" +
		"  const { pathname } = useLocation()\n" +
		"\n" +
		"  return (\n" +
		"    <div>\n" +
		"      <nav>\n" +
		"        <ul>\n" +
		"          <li><NavLink to=\"/dashboard\" end><LayoutDashboard /> 대시보드</NavLink></li>\n" +
		"          <li>\n" +
		"            <span>건물 관리</span>\n" +
		"            <ul>\n" +
		"              <li><NavLink to=\"/buildings\" end className={({ isActive }) => (isActive || pathname.startsWith('/buildings/') ? 'active' : undefined)}>건물 목록</NavLink></li>\n" +
		"            </ul>\n" +
		"          </li>\n" +
		"        </ul>\n" +
		"      </nav>\n" +
		"      <Breadcrumb />\n" +
		"      <Outlet />\n" +
		"    </div>\n" +
		"  )\n" +
		"}\n"
	if got != want {
		t.Errorf("got:\n%s\nwant:\n%s", got, want)
	}
	assertNotContains(t, got, "Legacy")
}
