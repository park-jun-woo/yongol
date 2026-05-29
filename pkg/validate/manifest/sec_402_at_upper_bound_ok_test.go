//ff:func feature=validate type=test control=sequence topic=manifest-auth
//ff:what SEC-402 테스트 — 경계값 30m 는 통과

package manifest

import "testing"

func TestSEC402_AtUpperBound_OK(t *testing.T) {
	if diags := sec402AccessTTLUpperBound(fsWithAccessTTL("30m")); len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics at boundary 30m, got %d", len(diags))
	}
}
