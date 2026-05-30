//ff:func feature=manifest type=test control=iteration dimension=1
//ff:what stripParamAndComma — 후행 콤마/파라미터 리스트 제거

package ddl

import "testing"

func TestStripParamAndComma(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"VARYING(255),", "VARYING"},
		{"VARYING(255)", "VARYING"},
		{"INT,", "INT"},
		{"INT", "INT"},
		{"(leading)", "(leading)"}, // idx==0 → not stripped
	}
	for _, c := range cases {
		if got := stripParamAndComma(c.in); got != c.want {
			t.Errorf("stripParamAndComma(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
