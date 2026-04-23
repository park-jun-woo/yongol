//ff:func feature=openapi-parse type=test control=sequence
//ff:what DeriveSuccessStatus — POST 가 200 만 선언 시 200 반환

package openapi

import "testing"

// TestDeriveSuccessStatus_PostFallsBackTo200 — POST operation that only
// declares 200 still gets a valid status. Common for idempotent POSTs.
func TestDeriveSuccessStatus_PostFallsBackTo200(t *testing.T) {
	got := DeriveSuccessStatus(opWith("200"), "POST")
	if got != 200 {
		t.Fatalf("POST 200 → %d, want 200", got)
	}
}
