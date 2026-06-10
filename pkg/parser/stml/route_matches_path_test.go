//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what TestRouteMatchesPath — 정확 일치·:param 세그먼트·말미 optional 생략/포함·불일치 케이스 검증

package stml

import "testing"

func TestRouteMatchesPath(t *testing.T) {
	cases := []struct {
		pattern, path string
		want          bool
	}{
		{"/login", "/login", true},
		{"/", "/", true},
		{"/workflows/:id", "/workflows/3", true},
		{"/workflows/:id", "/workflows", false},
		{"/workflows", "/login", false},
		{"/workflows/:id", "/workflows/3/edit", false},
		// trailing optional segments (":Name?") may be omitted or filled
		{"/buildings/:BuildingID/:PhotoID?", "/buildings/3", true},
		{"/buildings/:BuildingID/:PhotoID?", "/buildings/3/9", true},
		{"/buildings/:BuildingID/:PhotoID?", "/buildings", false},
		{"/unit-info/:BuildingID/:UnitID/:PhotoID?", "/unit-info/3/17", true},
		{"/webhooks/:id?", "/webhooks", true},
		{"/webhooks/:id?", "/webhooks/5", true},
		{"/webhooks/:id?", "/webhooks/5/extra", false},
		// a required segment is never optional
		{"/workflows/:id/:rev?", "/workflows", false},
	}
	for _, c := range cases {
		if got := RouteMatchesPath(c.pattern, c.path); got != c.want {
			t.Errorf("RouteMatchesPath(%q, %q) = %v, want %v", c.pattern, c.path, got, c.want)
		}
	}
}
