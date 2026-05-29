//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what integerGoType 단위 테스트

package ssac

import "testing"

func TestIntegerGoType(t *testing.T) {
	cases := map[string]string{
		"int64":   "integer64",
		"int32":   "integer32",
		"":        "integer",
		"unknown": "integer",
	}
	for in, want := range cases {
		if got := integerGoType(in); got != want {
			t.Errorf("integerGoType(%q) = %q, want %q", in, got, want)
		}
	}
}
