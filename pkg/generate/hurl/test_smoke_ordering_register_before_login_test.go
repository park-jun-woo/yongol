//ff:func feature=gen-hurl type=test control=sequence
//ff:what TestSmokeOrderingRegisterBeforeLogin — Register step이 Login step 보다 먼저 배치되는지 검증

package hurl

import (
	"testing"
)

// TestSmokeOrderingRegisterBeforeLogin pins BUG-015 Phase003: whenever
// the OpenAPI exposes both a Register and a Login operation, Register
// must appear first in smoke so an empty DB hits `POST /auth/register`
// (201) before `POST /auth/login` runs against a freshly-created user.
func TestSmokeOrderingRegisterBeforeLogin(t *testing.T) {
	fs := newSmokeFullstack(newAuthOnlyOpenAPI())
	steps := buildScenarioOrder(fs)
	opIDs := stepOpIDs(steps)
	regIdx := indexOfString(opIDs, "Register")
	loginIdx := indexOfString(opIDs, "Login")
	if regIdx < 0 {
		t.Fatalf("Register missing from smoke steps: %v", opIDs)
	}
	if loginIdx < 0 {
		t.Fatalf("Login missing from smoke steps: %v", opIDs)
	}
	if regIdx > loginIdx {
		t.Errorf("Register (idx=%d) must precede Login (idx=%d); got order=%v", regIdx, loginIdx, opIDs)
	}
}
