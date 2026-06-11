//ff:func feature=gen-react type=test control=sequence
//ff:what renderLayoutTSX sitemap 메뉴 — 전체 TSX 스냅샷 (NavLink·useLocation·lucide import + pathname + 메뉴/Outlet) 검증

package react

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderLayoutTSX_SitemapMenu(t *testing.T) {
	layout := stml.LayoutSpec{
		Name:      "app",
		HasOutlet: true,
		// surviving data-nav must be ignored on the sitemap path (TM-44
		// rejects it upstream; the emitter never mixes the two truths)
		NavItems: []stml.NavItem{{Path: "/legacy", Label: "Legacy"}},
	}
	menu := &sitemapMenu{Items: []sitemapMenuItem{
		{Kind: "page", Label: "대시보드", To: "/dashboard", Icon: "LayoutDashboard"},
		{Kind: "group", Label: "건물 관리", Children: []sitemapMenuItem{
			{Kind: "page", Label: "건물 목록", To: "/buildings", Prefixes: []string{"/buildings/"}},
		}},
	}}

	t.Run("full snapshot with icons and ancestor highlight", func(t *testing.T) {
		assertSitemapLayoutSnapshot(t, layout, menu)
	})

	t.Run("no prefixes and no icons trims the imports", func(t *testing.T) {
		plain := &sitemapMenu{Items: []sitemapMenuItem{
			{Kind: "page", Label: "대시보드", To: "/dashboard"},
		}}
		got := renderLayoutTSX("AppLayout", stml.LayoutSpec{Name: "app", HasOutlet: true}, nil, "", plain)
		assertContains(t, got, "import { NavLink, Outlet } from 'react-router-dom'")
		assertNotContains(t, got, "useLocation")
		assertNotContains(t, got, "lucide-react")
	})
}
