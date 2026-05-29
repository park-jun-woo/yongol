//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what isLiteral 단위 테스트 (Go 리터럴 판정)

package ssac

import "testing"

func TestIsLiteral(t *testing.T) {
	cases := map[string]bool{
		"true":     true,
		"false":    true,
		"nil":      true,
		`"hello"`:  true,
		"42":       true,
		"-7":       true,
		"x.Field":  false,
		"":         false,
		"variable": false,
	}
	for in, want := range cases {
		if got := isLiteral(in); got != want {
			t.Errorf("isLiteral(%q) = %v, want %v", in, got, want)
		}
	}
}
