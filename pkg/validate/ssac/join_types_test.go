//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-structural
//ff:what joinTypes — 타입 문자열 조인 검증 (empty/single/multiple)

package ssac

import "testing"

func TestJoinTypes(t *testing.T) {
	cases := []struct {
		name  string
		types []string
		want  string
	}{
		{name: "empty", types: nil, want: ""},
		{name: "single", types: []string{"bool"}, want: "bool"},
		{name: "two", types: []string{"FooResponse", "error"}, want: "FooResponse, error"},
		{name: "three", types: []string{"a", "b", "c"}, want: "a, b, c"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := joinTypes(c.types)
			if got != c.want {
				t.Errorf("joinTypes(%v) = %q, want %q", c.types, got, c.want)
			}
		})
	}
}
