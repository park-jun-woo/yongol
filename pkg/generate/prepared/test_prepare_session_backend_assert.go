//ff:func feature=generate type=test-helper control=sequence
//ff:what assertPrepareSessionBackend — prepareSessionBackend 반환값 단언

package prepared

import "testing"

// assertPrepareSessionBackend checks the Session returned by
// sessionBackendFor against a table entry's expectations.
func assertPrepareSessionBackend(t *testing.T, tc prepareSessionBackendCase, got *Session) {
	t.Helper()
	if tc.wantNil {
		if got != nil {
			t.Fatalf("expected nil, got %+v", got)
		}
		return
	}
	if got == nil {
		t.Fatalf("expected non-nil Session")
	}
	if got.Backend != tc.wantBE {
		t.Errorf("Backend=%q, want %q", got.Backend, tc.wantBE)
	}
}
