//ff:func feature=gen-react type=test control=sequence
//ff:what 단순 react 헬퍼 (pick*/sorted*/orDefault/quotedUnion/naivePluralize) 묶음 커버 — 같은 basename 미사용 함수용
package react

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestOrDefaultZC(t *testing.T) {
	if got := orDefault(nil, func(*manifest.FrontendTheme) string { return "x" }, "def"); got != "def" {
		t.Errorf("nil theme = %q", got)
	}
	theme := &manifest.FrontendTheme{Primary: "blue"}
	if got := orDefault(theme, pickPrimary, "def"); got != "blue" {
		t.Errorf("set value = %q", got)
	}
	if got := orDefault(theme, pickAccent, "def"); got != "def" {
		t.Errorf("empty value = %q", got)
	}
}
