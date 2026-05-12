//ff:func feature=gen-react type=test control=sequence
//ff:what layoutImports Link + Outlet import 검증

package react

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestLayoutImports_LinkAndOutlet(t *testing.T) {
	layout := stml.LayoutSpec{
		NavItems:  []stml.NavItem{{Path: "/foo", Label: "Foo"}},
		HasOutlet: true,
	}
	imports := layoutImports(layout)
	if len(imports) != 2 || imports[0] != "Link" || imports[1] != "Outlet" {
		t.Errorf("expected [Link, Outlet], got %v", imports)
	}
}
