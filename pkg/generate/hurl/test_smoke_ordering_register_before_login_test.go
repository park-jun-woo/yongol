//ff:func feature=gen-hurl type=test control=iteration dimension=1
//ff:what TestSmokeOrderingSignupBeforeLogin — 명명과 무관하게 signup-shape op 이 login-shape 보다 먼저 배치되는지 검증 (BUG-015 + BUG-023)

package hurl

import (
	"testing"
)

// TestSmokeOrderingSignupBeforeLogin pins BUG-015 Phase003 + BUG-023:
// whenever OpenAPI exposes both a signup-shape and a login-shape op,
// signup must appear first in smoke so an empty DB hits POST /auth/...
// (201) before the login runs against a freshly-created user. The
// previous implementation compared operationId literally against
// "Register" — any other name (Signup, Join, ...) silently broke the
// pair. Shape detection replaces that.
func TestSmokeOrderingSignupBeforeLogin(t *testing.T) {
	cases := []struct {
		name       string
		signupOpID string
		loginOpID  string
	}{
		{"Register+Login", "Register", "Login"},
		{"Signup+Login", "Signup", "Login"},
		{"Join+SignIn", "Join", "SignIn"},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assertSignupBeforeLogin(t, tc.signupOpID, tc.loginOpID)
		})
	}
}
