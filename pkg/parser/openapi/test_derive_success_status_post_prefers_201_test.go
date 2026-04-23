//ff:func feature=openapi-parse type=test control=sequence
//ff:what DeriveSuccessStatus — POST 가 201+200 중 201 선택

package openapi

import "testing"

// TestDeriveSuccessStatus_PostPrefers201 — POST with both 201 and 200
// picks 201 per RFC 9110 resource-creation convention.
func TestDeriveSuccessStatus_PostPrefers201(t *testing.T) {
	got := DeriveSuccessStatus(opWith("200", "201"), "POST")
	if got != 201 {
		t.Fatalf("POST 200+201 → %d, want 201", got)
	}
}
