//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what lcFirst 단위 테스트

package ssac

import "testing"

func TestLcFirst(t *testing.T) {
	cases := map[string]string{
		"":      "",
		"Hello": "hello",
		"hello": "hello",
		"ID":    "iD",
		"A":     "a",
		"9abc":  "9abc",
		"_x":    "_x",
	}
	for in, want := range cases {
		if got := lcFirst(in); got != want {
			t.Errorf("lcFirst(%q) = %q, want %q", in, got, want)
		}
	}
}
