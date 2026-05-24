//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-structural
//ff:what lcFirstEval — PascalCase -> camelCase 변환 검증 (정상/이미 소문자/빈 문자열)

package ssac

import "testing"

func TestLcFirstEval(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{name: "PascalCase", in: "IsZeroBalance", want: "isZeroBalance"},
		{name: "already lowercase", in: "isReady", want: "isReady"},
		{name: "single char", in: "A", want: "a"},
		{name: "empty", in: "", want: ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := lcFirstEval(c.in)
			if got != c.want {
				t.Errorf("lcFirstEval(%q) = %q, want %q", c.in, got, c.want)
			}
		})
	}
}
