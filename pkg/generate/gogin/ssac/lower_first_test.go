//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what lowerFirst 단위 테스트

package ssac

import "testing"

func TestLowerFirst(t *testing.T) {
	cases := map[string]string{
		"":                "",
		"PayloadTemplate": "payloadTemplate",
		"X":               "x",
		"already":         "already",
		"ID":              "iD",
	}
	for in, want := range cases {
		if got := lowerFirst(in); got != want {
			t.Errorf("lowerFirst(%q) = %q, want %q", in, got, want)
		}
	}
}
