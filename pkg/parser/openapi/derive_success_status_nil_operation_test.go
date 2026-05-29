//ff:func feature=openapi-parse type=test control=sequence
//ff:what DeriveSuccessStatus — nil operation 은 0 반환 (panic 금지)

package openapi

import "testing"

// TestDeriveSuccessStatus_NilOperation — defensive: nil op returns 0
// without panicking.
func TestDeriveSuccessStatus_NilOperation(t *testing.T) {
	if got := DeriveSuccessStatus(nil, "GET"); got != 0 {
		t.Fatalf("nil op → %d, want 0", got)
	}
}
