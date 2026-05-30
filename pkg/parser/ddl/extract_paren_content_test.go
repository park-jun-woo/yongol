//ff:func feature=manifest type=test control=iteration dimension=1
//ff:what extractParenContent — 괄호 내부 추출 / 누락·역순 괄호 빈 문자열

package ddl

import "testing"

func TestExtractParenContent(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"foo (a, b)", "a, b"},
		{"(  trimmed  )", "trimmed"},
		{"no parens", ""},
		{"only open (", ""},
		{"only close )", ""},
		{")reversed(", ""},
		{"()", ""},
	}
	for _, c := range cases {
		if got := extractParenContent(c.in); got != c.want {
			t.Errorf("extractParenContent(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
