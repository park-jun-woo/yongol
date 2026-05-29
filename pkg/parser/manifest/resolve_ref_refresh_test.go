//ff:func feature=manifest type=test control=sequence
//ff:what ResolveRef_RefreshTokenTTL — 168h → 604800초 변환 검증

package manifest

import "testing"

func TestResolveRef_RefreshTokenTTL(t *testing.T) {
	cfg := &ProjectConfig{
		Backend: Backend{
			Auth: &Auth{RefreshTokenTTL: "168h"},
		},
	}
	rv, ok := ResolveRef(cfg, "auth.refreshTokenTTL")
	if !ok {
		t.Fatal("expected ok=true")
	}
	if rv.GoLit != "604800" {
		t.Errorf("GoLit = %q, want %q", rv.GoLit, "604800")
	}
}
