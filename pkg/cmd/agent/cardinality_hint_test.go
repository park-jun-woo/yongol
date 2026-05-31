//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestCardinalityHint — op/path 로부터 sqlc cardinality 어노테이션 추론 검증
package agent

import (
	"testing"
)

func TestCardinalityHint(t *testing.T) {
	cases := []struct {
		op   string
		path string
		want string
	}{
		{"getUser", "/users/{id}", ":one"},
		{"listUsers", "/users", ":many"},
		{"createUser", "/users", ":one RETURNING"},
		{"updateUser", "/users/{id}", ":one RETURNING"},
		{"deleteUser", "/users/{id}", ":exec"},
		{"removeUser", "/users/{id}", ":exec"},
		{"doSomethingWeird", "/x", ":one"},
	}
	for _, c := range cases {
		if got := cardinalityHint(c.op, c.path); got != c.want {
			t.Errorf("cardinalityHint(%q,%q) = %q, want %q", c.op, c.path, got, c.want)
		}
	}
}
