//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestStringSliceEqual — 길이/요소 동등 비교
package migration

import (
	"testing"
)

func TestStringSliceEqual(t *testing.T) {
	cases := []struct {
		name string
		a, b []string
		want bool
	}{
		{"equal", []string{"a", "b"}, []string{"a", "b"}, true},
		{"len diff", []string{"a"}, []string{"a", "b"}, false},
		{"elem diff", []string{"a", "x"}, []string{"a", "b"}, false},
		{"both empty", nil, []string{}, true},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if got := stringSliceEqual(c.a, c.b); got != c.want {
				t.Errorf("stringSliceEqual(%v, %v) = %v, want %v", c.a, c.b, got, c.want)
			}
		})
	}
}
