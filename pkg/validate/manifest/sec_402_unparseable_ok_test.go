//ff:func feature=validate type=test control=sequence topic=manifest-auth
//ff:what SEC-402 테스트 — 파싱 불가 값은 통과 (다른 규칙이 포맷 검증 담당)

package manifest

import "testing"

func TestSEC402_Unparseable_OK(t *testing.T) {
	if diags := sec402AccessTTLUpperBound(fsWithAccessTTL("nonsense")); len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics when unparseable, got %d", len(diags))
	}
}
