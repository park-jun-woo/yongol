//ff:func feature=manifest type=test control=sequence
//ff:what ResolveRef_NilAuth — nil auth 에서 false 반환 검증

package manifest

import "testing"

func TestResolveRef_NilAuth(t *testing.T) {
	cfg := &ProjectConfig{}
	_, ok := ResolveRef(cfg, "auth.accessTokenTTL")
	if ok {
		t.Error("expected ok=false for nil auth")
	}
}
