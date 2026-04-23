//ff:func feature=openapi-parse type=test control=sequence
//ff:what DeriveSuccessStatus — DELETE 가 200+204 중 204 선택

package openapi

import "testing"

// TestDeriveSuccessStatus_DeletePrefers204 — DELETE with 204 + 200
// picks 204 (No Content is the conventional success for DELETE).
func TestDeriveSuccessStatus_DeletePrefers204(t *testing.T) {
	got := DeriveSuccessStatus(opWith("200", "204"), "DELETE")
	if got != 204 {
		t.Fatalf("DELETE 200+204 → %d, want 204", got)
	}
}
