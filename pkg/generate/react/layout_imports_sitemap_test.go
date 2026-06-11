//ff:func feature=gen-react type=test control=sequence
//ff:what layoutImports sitemap 분기 — NavLink/useLocation 게이트, data-nav Link 미포함, 빈 메뉴 검증

package react

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestLayoutImports_SitemapMenu(t *testing.T) {
	layout := stml.LayoutSpec{
		HasOutlet: true,
		NavItems:  []stml.NavItem{{Path: "/legacy", Label: "Legacy"}}, // ignored on the sitemap path
	}

	t.Run("full set", func(t *testing.T) {
		menu := &sitemapMenu{Items: []sitemapMenuItem{
			{Kind: "page", To: "/a", Prefixes: []string{"/a/"}},
		}}
		got := layoutImports(layout, true, menu)
		want := []string{"NavLink", "Outlet", "useLocation", "useNavigate"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("imports = %v, want %v", got, want)
		}
	})

	t.Run("empty sitemap menu needs no Link at all", func(t *testing.T) {
		got := layoutImports(layout, false, &sitemapMenu{})
		want := []string{"Outlet"}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("imports = %v, want %v", got, want)
		}
	})
}
