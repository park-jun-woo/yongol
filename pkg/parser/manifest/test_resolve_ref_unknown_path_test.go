//ff:func feature=manifest type=test control=sequence
//ff:what ResolveRef_UnknownPath — 미지원 경로에서 false 반환 검증

package manifest

import "testing"

func TestResolveRef_UnknownPath(t *testing.T) {
	cfg := &ProjectConfig{
		Backend: Backend{
			Auth: &Auth{AccessTokenTTL: "15m"},
		},
	}
	_, ok := ResolveRef(cfg, "auth.unknownField")
	if ok {
		t.Error("expected ok=false for unknown path")
	}
}
