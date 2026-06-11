//ff:func feature=gen-react type=test control=iteration dimension=1
//ff:what staticRoutePrefix — 파라미터 절단/정적 패턴 그대로/선행 파라미터·빈 패턴 경계 검증

package react

import "testing"

func TestStaticRoutePrefix(t *testing.T) {
	cases := []struct{ in, want string }{
		{"/buildings/:BuildingID", "/buildings/"},
		{"/a/:B/c", "/a/"},
		{"/building-documents", "/building-documents"},
		{"/:OrgID/buildings", "/"},
		{"", ""},
		{"/", "/"},
	}
	for _, c := range cases {
		if got := staticRoutePrefix(c.in); got != c.want {
			t.Errorf("staticRoutePrefix(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
