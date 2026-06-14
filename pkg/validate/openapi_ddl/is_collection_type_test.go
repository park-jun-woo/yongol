//ff:func feature=validate type=test control=iteration dimension=1 topic=openapi-ddl
//ff:what isCollectionType — []T/Page[T]/Cursor[T] 컬렉션 판별, 단일/빈/0인덱스 괄호 경계

package openapi_ddl

import "testing"

func TestIsCollectionType(t *testing.T) {
	cases := []struct {
		raw  string
		want bool
	}{
		{"[]Gig", true},
		{"Page[Gig]", true},
		{"Cursor[Gig]", true},
		{"*User", false},
		{"User", false},
		{"billing.User", false},
		{"", false},
		{"[", false}, // bracket at index 0, not >0
	}
	for _, c := range cases {
		if got := isCollectionType(c.raw); got != c.want {
			t.Errorf("isCollectionType(%q) = %v, want %v", c.raw, got, c.want)
		}
	}
}
