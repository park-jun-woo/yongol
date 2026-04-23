//ff:func feature=validate type=test control=sequence topic=manifest-auth
//ff:what SEC-402 테스트 — 15m 는 상한 이하로 통과

package manifest

import "testing"

func TestSEC402_BelowUpperBound_OK(t *testing.T) {
	if diags := sec402AccessTTLUpperBound(fsWithAccessTTL("15m")); len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics at 15m, got %d", len(diags))
	}
}
