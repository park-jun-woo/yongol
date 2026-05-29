//ff:func feature=chain type=test control=iteration dimension=1
//ff:what ownershipIcon 가 ownership 종류별 아이콘 문자열을 반환하는지 검증
package chain

import "testing"

func TestOwnershipIcon(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"preserve", "preserve"},
		{"gen", "gen"},
		{"", ""},
		{"unknown", ""},
	}
	for _, c := range cases {
		if got := ownershipIcon(c.in); got != c.want {
			t.Errorf("ownershipIcon(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
