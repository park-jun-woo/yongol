//ff:func feature=agent type=test control=sequence
//ff:what TestDomainFromPath — API path 에서 도메인 추출(복수형 단수화, kebab→snake) 검증

package agent

import "testing"

func TestDomainFromPath(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"/workflows/{id}", "workflow"},
		{"/auth/login", "auth"},
		{"/payment-intents/{id}", "payment_intent"},
		{"GET /users", "user"},
		{"/", "default"},
		{"", "default"},
		{"/address", "address"}, // "ss" suffix is preserved (not stripped)
	}
	for _, c := range cases {
		if got := domainFromPath(c.in); got != c.want {
			t.Errorf("domainFromPath(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
