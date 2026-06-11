//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what TestFirstRequiredSegment — 필수 :Name 검출, 선택 :Name? 건너뜀, 없음 → "" 검증

package stml_openapi

import "testing"

func TestFirstRequiredSegment(t *testing.T) {
	cases := []struct {
		pattern string
		want    string
	}{
		{"/buildings/:BuildingID", ":BuildingID"},
		{"/buildings/:BuildingID/:RoomID", ":BuildingID"},
		{"/members/:Page?", ""},
		{"/members/:Page?/:MemberID", ":MemberID"},
		{"/login", ""},
		{"/", ""},
		{"", ""},
	}
	for _, c := range cases {
		if got := firstRequiredSegment(c.pattern); got != c.want {
			t.Errorf("firstRequiredSegment(%q) = %q, want %q", c.pattern, got, c.want)
		}
	}
}
