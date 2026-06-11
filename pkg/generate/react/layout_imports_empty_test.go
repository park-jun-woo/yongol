//ff:func feature=gen-react type=test control=sequence
//ff:what layoutImports 빈 import 검증

package react

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestLayoutImports_Empty(t *testing.T) {
	layout := stml.LayoutSpec{}
	imports := layoutImports(layout, false, nil)
	if len(imports) != 0 {
		t.Errorf("expected empty imports, got %v", imports)
	}
}
