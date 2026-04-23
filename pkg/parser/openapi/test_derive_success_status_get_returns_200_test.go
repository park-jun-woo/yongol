//ff:func feature=openapi-parse type=test control=sequence
//ff:what DeriveSuccessStatus — GET + 200 단일 선언

package openapi

import "testing"

// TestDeriveSuccessStatus_GetReturns200 — GET operation with only the
// default 200 declaration.
func TestDeriveSuccessStatus_GetReturns200(t *testing.T) {
	got := DeriveSuccessStatus(opWith("200"), "GET")
	if got != 200 {
		t.Fatalf("GET 200 → %d, want 200", got)
	}
}
