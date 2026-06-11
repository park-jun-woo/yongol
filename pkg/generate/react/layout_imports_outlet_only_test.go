//ff:func feature=gen-react type=test control=sequence
//ff:what layoutImports Outlet only import 검증

package react

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestLayoutImports_OutletOnly(t *testing.T) {
	layout := stml.LayoutSpec{HasOutlet: true}
	imports := layoutImports(layout, false, nil)
	if len(imports) != 1 || imports[0] != "Outlet" {
		t.Errorf("expected [Outlet], got %v", imports)
	}
}
