//ff:func feature=gen-react type=test control=iteration dimension=1
//ff:what layoutImports — logout 방출 시 useNavigate import 추가 검증

package react

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestLayoutImports_Logout(t *testing.T) {
	layout := stml.LayoutSpec{
		NavItems:  []stml.NavItem{{Path: "dashboard", Label: "대시보드"}},
		HasOutlet: true,
		Logout:    &stml.LogoutSpec{OperationID: "Logout"},
	}
	imports := layoutImports(layout, true)
	want := []string{"Link", "Outlet", "useNavigate"}
	if len(imports) != len(want) {
		t.Fatalf("imports = %v, want %v", imports, want)
	}
	for i := range want {
		if imports[i] != want[i] {
			t.Errorf("imports[%d] = %q, want %q", i, imports[i], want[i])
		}
	}

	// emission gated off (no auth): no useNavigate even with a declared logout.
	imports = layoutImports(layout, false)
	if len(imports) != 2 || imports[0] != "Link" || imports[1] != "Outlet" {
		t.Errorf("expected [Link, Outlet] without logout emission, got %v", imports)
	}
}
