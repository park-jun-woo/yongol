//ff:func feature=validate type=test control=sequence topic=manifest-auth
//ff:what SEC-402 테스트 — 빈 값은 통과

package manifest

import "testing"

func TestSEC402_Empty_OK(t *testing.T) {
	if diags := sec402AccessTTLUpperBound(fsWithAccessTTL("")); len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics when empty, got %d", len(diags))
	}
}
