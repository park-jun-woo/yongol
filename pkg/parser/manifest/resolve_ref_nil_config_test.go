//ff:func feature=manifest type=test control=sequence
//ff:what ResolveRef_NilConfig — nil config 에서 false 반환 검증

package manifest

import "testing"

func TestResolveRef_NilConfig(t *testing.T) {
	_, ok := ResolveRef(nil, "auth.accessTokenTTL")
	if ok {
		t.Error("expected ok=false for nil config")
	}
}
