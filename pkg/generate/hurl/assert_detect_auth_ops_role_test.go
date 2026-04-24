//ff:func feature=gen-hurl type=test-helper control=selection
//ff:what assertDetectAuthOpsRole — detectAuthOps 결과를 기대 role 과 비교

package hurl

import (
	"testing"
)

// assertDetectAuthOpsRole compares the detected role against the fixture's
// expectation using the opID-match convention.
func assertDetectAuthOpsRole(t *testing.T, tc detectAuthOpsFixture, signup, login *detectedAuthOp, warns []string) {
	t.Helper()
	got := ""
	switch {
	case signup != nil && signup.OpID == tc.opID:
		got = "signup"
	case login != nil && login.OpID == tc.opID:
		got = "login"
	}
	if got != tc.wantRole {
		t.Errorf("role: want %q got %q (warns=%v)", tc.wantRole, got, warns)
	}
}
