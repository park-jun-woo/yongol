//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what firstLine 단위 테스트

package ssac

import "testing"

func TestFirstLine(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"single line", "hello", "hello"},
		{"leading blank lines", "\n\n  first\nsecond", "first"},
		{"trims whitespace", "   only   ", "only"},
		{"empty", "", ""},
		{"all blank", "\n  \n\t", ""},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := firstLine(tc.in); got != tc.want {
				t.Errorf("firstLine(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}
