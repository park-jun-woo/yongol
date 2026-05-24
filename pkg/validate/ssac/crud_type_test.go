//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-structural
//ff:what crudType — CRUD 타입 판정 (get/post/put/delete/non-CRUD) 검증

package ssac

import (
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestCrudType(t *testing.T) {
	cases := []struct {
		typ  string
		want bool
	}{
		{"get", true},
		{"post", true},
		{"put", true},
		{"delete", true},
		{"empty", false},
		{"exists", false},
		{"state", false},
		{"auth", false},
		{"call", false},
		{"response", false},
		{"", false},
	}
	for _, c := range cases {
		t.Run(c.typ, func(t *testing.T) {
			seq := parsessac.Sequence{Type: c.typ}
			got := crudType(seq)
			if got != c.want {
				t.Errorf("crudType(%q) = %v, want %v", c.typ, got, c.want)
			}
		})
	}
}
