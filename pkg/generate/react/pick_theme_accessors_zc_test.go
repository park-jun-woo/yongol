//ff:func feature=gen-react type=test control=sequence
//ff:what 단순 react 헬퍼 (pick*/sorted*/orDefault/quotedUnion/naivePluralize) 묶음 커버 — 같은 basename 미사용 함수용
package react

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestPickThemeAccessorsZC(t *testing.T) {
	th := &manifest.FrontendTheme{
		Primary: "p", Secondary: "s", Accent: "a", Destructive: "d",
		Muted: "m", Background: "bg", Foreground: "fg", Border: "bd", Radius: "0.5rem",
	}
	if pickPrimary(th) != "p" || pickSecondary(th) != "s" || pickAccent(th) != "a" ||
		pickDestructive(th) != "d" || pickMuted(th) != "m" || pickBackground(th) != "bg" ||
		pickForeground(th) != "fg" || pickBorder(th) != "bd" || pickRadius(th) != "0.5rem" {
		t.Errorf("pick accessor mismatch")
	}
}
