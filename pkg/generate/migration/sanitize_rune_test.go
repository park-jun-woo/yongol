//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestSanitizeRune — a-z/0-9/_ 유지, 그 외는 _
package migration

import (
	"testing"
)

func TestSanitizeRune(t *testing.T) {
	cases := []struct {
		in   rune
		want rune
	}{
		{'a', 'a'},
		{'5', '5'},
		{'_', '_'},
		{'-', '_'},
		{'A', '_'},
		{'.', '_'},
	}
	for _, c := range cases {
		if got := sanitizeRune(c.in); got != c.want {
			t.Errorf("sanitizeRune(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
