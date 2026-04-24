//ff:func feature=gen-hurl type=test-helper control=sequence
//ff:what assertSignupBeforeLogin — smoke 내 signup 이 login 보다 앞에 위치하는지 검증

package hurl

import (
	"testing"
)

// assertSignupBeforeLogin asserts the smoke scenario ordering invariant:
// the signup op must appear before the login op. Fails the test with a
// descriptive diagnostic on mismatch.
func assertSignupBeforeLogin(t *testing.T, signupOpID, loginOpID string) {
	t.Helper()
	doc := synthAuthDoc(signupOpID, loginOpID)
	fs := newSmokeFullstack(doc)
	steps := buildScenarioOrder(fs)
	opIDs := stepOpIDs(steps)
	sIdx := indexOfString(opIDs, signupOpID)
	lIdx := indexOfString(opIDs, loginOpID)
	if sIdx < 0 {
		t.Fatalf("%s missing from smoke steps: %v", signupOpID, opIDs)
	}
	if lIdx < 0 {
		t.Fatalf("%s missing from smoke steps: %v", loginOpID, opIDs)
	}
	if sIdx > lIdx {
		t.Errorf("%s (idx=%d) must precede %s (idx=%d); got order=%v",
			signupOpID, sIdx, loginOpID, lIdx, opIDs)
	}
}
