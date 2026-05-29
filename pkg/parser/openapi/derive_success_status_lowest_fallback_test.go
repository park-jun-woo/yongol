//ff:func feature=openapi-parse type=test control=sequence
//ff:what DeriveSuccessStatus — 우선순위에 없는 2xx 면 최저 코드 선택

package openapi

import "testing"

// TestDeriveSuccessStatus_LowestFallback — odd case: declared 2xx is not
// in the method's preference list (e.g. 207 Multi-Status). Fallback
// picks the lowest-numbered declared 2xx deterministically.
func TestDeriveSuccessStatus_LowestFallback(t *testing.T) {
	got := DeriveSuccessStatus(opWith("207", "208"), "POST")
	if got != 207 {
		t.Fatalf("POST 207+208 → %d, want 207", got)
	}
}
