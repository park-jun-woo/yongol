//ff:func feature=openapi-parse type=test control=sequence
//ff:what DeriveSuccessStatus — 2xx 가 없으면 0 반환

package openapi

import "testing"

// TestDeriveSuccessStatus_NoTwoXX — operation declares only 4xx responses
// → function returns 0 so callers surface a XOS-22/XOS-81 diagnostic.
func TestDeriveSuccessStatus_NoTwoXX(t *testing.T) {
	got := DeriveSuccessStatus(opWith("400"), "POST")
	if got != 0 {
		t.Fatalf("POST 400 → %d, want 0", got)
	}
}
