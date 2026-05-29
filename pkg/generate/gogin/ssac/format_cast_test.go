//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what formatPrimitiveCast 단위 테스트
package ssac

import "testing"

func TestFormatPrimitiveCast(t *testing.T) {
	cases := map[string]string{
		"email":     "string",
		"uuid":      "",
		"":          "",
		"int64":     "",
		"date-time": "",
		"unknown":   "",
	}
	for in, want := range cases {
		if got := formatPrimitiveCast(in); got != want {
			t.Errorf("formatPrimitiveCast(%q) = %q, want %q", in, got, want)
		}
	}
}
