//ff:func feature=manifest type=test control=sequence
//ff:what ResolveRef_AccessTokenTTL — 15m → 900초 변환 검증

package manifest

import "testing"

func TestResolveRef_AccessTokenTTL(t *testing.T) {
	cfg := &ProjectConfig{
		Backend: Backend{
			Auth: &Auth{AccessTokenTTL: "15m"},
		},
	}
	rv, ok := ResolveRef(cfg, "auth.accessTokenTTL")
	if !ok {
		t.Fatal("expected ok=true")
	}
	if rv.Raw != "15m" {
		t.Errorf("Raw = %q, want %q", rv.Raw, "15m")
	}
	if rv.GoLit != "900" {
		t.Errorf("GoLit = %q, want %q", rv.GoLit, "900")
	}
	if rv.GoType != "int64" {
		t.Errorf("GoType = %q, want %q", rv.GoType, "int64")
	}
}
