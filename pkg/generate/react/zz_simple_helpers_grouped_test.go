//ff:func feature=gen-react type=test control=sequence
//ff:what 단순 react 헬퍼 (pick*/sorted*/orDefault/quotedUnion/naivePluralize) 묶음 커버 — 같은 basename 미사용 함수용
package react

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestNaivePluralizeZC(t *testing.T) {
	cases := map[string]string{
		"": "", "course": "courses", "class": "classes",
		"dish": "dishes", "match": "matches", "box": "boxes", "quiz": "quizes",
	}
	for in, want := range cases {
		if got := naivePluralize(in); got != want {
			t.Errorf("naivePluralize(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestQuotedUnionZC(t *testing.T) {
	if got := quotedUnion([]string{"a", "b"}); got != "'a' | 'b'" {
		t.Errorf("quotedUnion = %q", got)
	}
	if got := quotedUnion(nil); got != "" {
		t.Errorf("quotedUnion(nil) = %q", got)
	}
}

func TestSortedKeysZC(t *testing.T) {
	if got := sortedKeys(map[string]string{"b": "1", "a": "2"}); !reflect.DeepEqual(got, []string{"a", "b"}) {
		t.Errorf("sortedKeys = %v", got)
	}
}

func TestSortedMapKeysZC(t *testing.T) {
	if got := sortedMapKeys(map[string]string{"y": "1", "x": "2"}); !reflect.DeepEqual(got, []string{"x", "y"}) {
		t.Errorf("sortedMapKeys = %v", got)
	}
}

func TestSortedLayoutNamesZC(t *testing.T) {
	grouped := map[string][]stmlRoute{"": nil, "app": nil, "auth": nil}
	got := sortedLayoutNames(grouped)
	if len(got) != 3 || got[0] != "app" || got[1] != "auth" || got[2] != "" {
		t.Errorf("sortedLayoutNames = %v", got)
	}
}

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
