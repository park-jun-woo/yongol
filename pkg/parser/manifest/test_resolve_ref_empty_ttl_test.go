//ff:func feature=manifest type=test control=sequence
//ff:what ResolveRef_EmptyTTL — 빈 TTL 에서 false 반환 검증

package manifest

import "testing"

func TestResolveRef_EmptyTTL(t *testing.T) {
	cfg := &ProjectConfig{
		Backend: Backend{
			Auth: &Auth{AccessTokenTTL: ""},
		},
	}
	_, ok := ResolveRef(cfg, "auth.accessTokenTTL")
	if ok {
		t.Error("expected ok=false for empty TTL")
	}
}
