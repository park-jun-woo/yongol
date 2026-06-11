//ff:func feature=gen-react type=test control=sequence
//ff:what layoutImports — DynamicCrumb 시 useLocation 포함 / Outlet 없으면 미포함 검증

package react

import (
	"slices"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestLayoutImportsDynamicCrumb(t *testing.T) {
	menu := &sitemapMenu{Items: []sitemapMenuItem{{Kind: "page", Label: "홈", To: "/"}}, DynamicCrumb: true}

	t.Run("DynamicCrumb with an Outlet needs useLocation for the reset", func(t *testing.T) {
		got := layoutImports(stml.LayoutSpec{Name: "app", HasOutlet: true}, false, menu)
		if !slices.Contains(got, "useLocation") {
			t.Errorf("imports = %v, want useLocation for the pathname reset", got)
		}
	})

	t.Run("DynamicCrumb without an Outlet adds nothing", func(t *testing.T) {
		got := layoutImports(stml.LayoutSpec{Name: "bare"}, false, menu)
		if slices.Contains(got, "useLocation") {
			t.Errorf("imports = %v, useLocation must not appear without an Outlet", got)
		}
	})
}
