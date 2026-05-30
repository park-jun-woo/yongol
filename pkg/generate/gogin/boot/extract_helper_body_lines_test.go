//ff:func feature=gen-gogin type=test control=sequence
//ff:what extractHelperBodyLines — top-level func 선언 문자열에서 body 라인 추출 (signature/braces 제거)

package boot

import (
	"strings"
	"testing"
)

func TestExtractHelperBodyLines(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want []string
	}{
		{"no open brace", "func f() int", nil},
		{"open brace without close", "func f() int {", nil},
		{
			"single line body",
			"func f() int { return 1 }",
			[]string{" return 1 "},
		},
		{
			"multiline body",
			"func f() int {\n\treturn 1\n}",
			[]string{"", "\treturn 1", ""},
		},
	}
	for _, c := range cases {
		got := extractHelperBodyLines(c.in)
		if strings.Join(got, "|") != strings.Join(c.want, "|") {
			t.Errorf("%s: extractHelperBodyLines = %#v, want %#v", c.name, got, c.want)
		}
	}
}
