//ff:func feature=cli-init type=test control=sequence
//ff:what TestWriteProjectIDBoundary — i==0/camelCase/acronym/무경계에서 언더스코어 삽입 여부 검증

package cliinit

import (
	"strings"
	"testing"
)

func boundaryAt(s string, i int) string {
	runes := []rune(s)
	var b strings.Builder
	writeProjectIDBoundary(&b, runes, i, runes[i])
	return b.String()
}

func TestWriteProjectIDBoundary(t *testing.T) {
	cases := []struct {
		name string
		s    string
		i    int
		want string // "_" means underscore emitted, "" means none
	}{
		{"index zero", "Abc", 0, ""},
		{"camel boundary lower->Upper", "fooBar", 3, "_"},  // 'B' after 'o'
		{"digit->Upper", "v2Api", 2, "_"},                  // 'A' after '2'
		{"acronym boundary", "HTTPServer", 4, "_"},         // 'S' is Upper, prev 'P' Upper, next 'e' lower
		{"no boundary upper->upper no lower next", "ABC", 2, ""}, // 'C' Upper, prev Upper, no next
		{"no boundary lower->lower", "abcd", 2, ""},
	}
	for _, c := range cases {
		if got := boundaryAt(c.s, c.i); got != c.want {
			t.Errorf("%s: boundaryAt(%q,%d) = %q, want %q", c.name, c.s, c.i, got, c.want)
		}
	}
}
