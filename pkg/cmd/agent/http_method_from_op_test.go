//ff:func feature=agent type=test control=selection
//ff:what TestHTTPMethodFromOp — operationId 접두사로 HTTP method 추론 검증

package agent

import "testing"

func TestHTTPMethodFromOp(t *testing.T) {
	cases := []struct {
		op   string
		want string
	}{
		{"listUsers", "get"},
		{"getUser", "get"},
		{"createUser", "post"},
		{"updateUser", "put"},
		{"deleteUser", "delete"},
		{"removeUser", "delete"},
		{"doThing", "post"},
	}
	for _, c := range cases {
		if got := httpMethodFromOp(c.op); got != c.want {
			t.Errorf("httpMethodFromOp(%q) = %q, want %q", c.op, got, c.want)
		}
	}
}
