//ff:func feature=gen-react type=test control=iteration dimension=1
//ff:what stripOptionalSegments — optional 세그먼트 제거 / 필수 세그먼트 보존 / 전부 optional → "/" 검증

package react

import "testing"

func TestStripOptionalSegments(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"/unit-list", "/unit-list"},                                  // no params
		{"/unit-list/:BuildingID?", "/unit-list"},                     // optional stripped
		{"/unit-info/:BuildingID/:UnitID?", "/unit-info/:BuildingID"}, // required kept (TM-34 blocks it upstream)
		{"/:All?", "/"}, // optional-only collapses to the index
		{"/", "/"},      // root unchanged
	}
	for _, c := range cases {
		if got := stripOptionalSegments(c.in); got != c.want {
			t.Errorf("stripOptionalSegments(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
