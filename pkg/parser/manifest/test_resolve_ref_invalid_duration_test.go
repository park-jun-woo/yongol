//ff:func feature=manifest type=test control=sequence
//ff:what ResolveRef_InvalidDuration — 잘못된 duration 에서 false 반환 검증

package manifest

import "testing"

func TestResolveRef_InvalidDuration(t *testing.T) {
	cfg := &ProjectConfig{
		Backend: Backend{
			Auth: &Auth{AccessTokenTTL: "notaduration"},
		},
	}
	_, ok := ResolveRef(cfg, "auth.accessTokenTTL")
	if ok {
		t.Error("expected ok=false for invalid duration")
	}
}
