//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestMatchFKAction — prefix 뒤 FK 동작 토큰 추출
package migration

import (
	"testing"
)

func TestMatchFKAction(t *testing.T) {
	cases := []struct {
		tail, prefix, want string
	}{
		{"ON DELETE CASCADE", "ON DELETE", "CASCADE"},
		{"ON DELETE SET NULL", "ON DELETE", "SET NULL"},
		{"ON UPDATE RESTRICT", "ON UPDATE", "RESTRICT"},
		{"ON DELETE NO ACTION", "ON DELETE", "NO ACTION"},
		{"ON DELETE CASCADE", "ON UPDATE", ""},
		{"REFERENCES x(id)", "ON DELETE", ""},
	}
	for _, c := range cases {
		if got := matchFKAction(c.tail, c.prefix); got != c.want {
			t.Errorf("matchFKAction(%q,%q) = %q, want %q", c.tail, c.prefix, got, c.want)
		}
	}
}
