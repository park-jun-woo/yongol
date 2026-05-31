//ff:func feature=gen-react type=test control=sequence
//ff:what TestByName_ZeroCov — react 코드젠 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package react

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestByNameLayoutTSX_ZeroCov(t *testing.T) {
	layout := stml.LayoutSpec{
		Name:      "app",
		NavItems:  []stml.NavItem{{Path: "/home", Label: "Home"}},
		HasOutlet: true,
	}
	out := renderLayoutTSX("AppLayout", layout)
	if !strings.Contains(out, "AppLayout") {
		t.Errorf("renderLayoutTSX missing name")
	}
}
