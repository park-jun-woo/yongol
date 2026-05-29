//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what isIntegerLiteralStr 단위 테스트

package ssac

import "testing"

func TestIsIntegerLiteralStr(t *testing.T) {
	cases := map[string]bool{
		"":    false,
		"0":   true,
		"123": true,
		"-7":  true,
		"-":   false,
		"12a": false,
		"a12": false,
		"1.5": false,
		"  3": false,
		"-0":  true,
	}
	for in, want := range cases {
		if got := isIntegerLiteralStr(in); got != want {
			t.Errorf("isIntegerLiteralStr(%q) = %v, want %v", in, got, want)
		}
	}
}
