//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-manifest
//ff:what isAuthPath — auth 경로 heuristic 매칭 검증

package hurl_manifest

import "testing"

func TestIsAuthPath(t *testing.T) {
	cases := []struct {
		name string
		path string
		want bool
	}{
		{name: "auth_segment", path: "/auth/login", want: true},
		{name: "auth_segment_register", path: "/api/auth/register", want: true},
		{name: "login_suffix", path: "/api/login", want: true},
		{name: "signin_suffix", path: "/api/signin", want: true},
		{name: "register_suffix", path: "/api/register", want: true},
		{name: "signup_suffix", path: "/api/signup", want: true},
		{name: "login_with_trailing", path: "/api/login/callback", want: true},
		{name: "uppercase_auth", path: "/AUTH/login", want: true},
		{name: "mixed_case_Login", path: "/api/Login", want: true},
		{name: "non_auth_path", path: "/api/users", want: false},
		{name: "non_auth_orders", path: "/api/orders/123", want: false},
		{name: "empty_path", path: "", want: false},
		{name: "partial_login_in_name", path: "/api/loginHistory", want: false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := isAuthPath(c.path)
			if got != c.want {
				t.Errorf("isAuthPath(%q) = %v, want %v", c.path, got, c.want)
			}
		})
	}
}
