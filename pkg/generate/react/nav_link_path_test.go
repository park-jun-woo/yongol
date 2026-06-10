//ff:func feature=gen-react type=test control=iteration dimension=1
//ff:what navLinkPath — 정적 경로 불변/페이지명 치환/optional·필수 세그먼트 strip/미지 페이지 폴백 검증

package react

import "testing"

func TestNavLinkPath(t *testing.T) {
	patterns := map[string]string{
		"dashboard": "/dashboard",
		"unit-list": "/unit-list/:UnitID?",
		"detail":    "/buildings/:BuildingID",
		"root":      "/",
	}

	cases := []struct {
		name, target, want string
	}{
		{"static path verbatim", "/anything/3", "/anything/3"},
		{"page name resolved", "dashboard", "/dashboard"},
		{"optional segment stripped", "unit-list", "/unit-list"},
		{"required segment stripped defensively", "detail", "/buildings"},
		{"root pattern", "root", "/"},
		{"unknown page falls back", "nowhere", "/nowhere"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := navLinkPath(c.target, patterns); got != c.want {
				t.Errorf("navLinkPath(%q) = %q, want %q", c.target, got, c.want)
			}
		})
	}
}
