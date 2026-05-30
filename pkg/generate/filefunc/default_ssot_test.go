//ff:func feature=gen-filefunc type=test control=sequence
//ff:what TestDefaultSSOTEntries — 고정 ssot 키 맵 반환 검증

package filefunc

import "testing"

func TestDefaultSSOTEntries(t *testing.T) {
	got := defaultSSOTEntries()
	want := []string{
		"openapi", "ddl", "ssac", "states", "policy",
		"scenario", "funcspec", "config",
	}
	if len(got) != len(want) {
		t.Errorf("expected %d ssot keys, got %d: %v", len(want), len(got), got)
	}
	for _, k := range want {
		if _, ok := got[k]; !ok {
			t.Errorf("missing ssot key %q: %v", k, got)
		}
	}
}
