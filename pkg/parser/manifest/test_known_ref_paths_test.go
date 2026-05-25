//ff:func feature=manifest type=test control=iteration dimension=1
//ff:what KnownRefPaths — 지원 경로 목록에 필수 항목 포함 검증

package manifest

import "testing"

func TestKnownRefPaths(t *testing.T) {
	paths := KnownRefPaths()
	if len(paths) == 0 {
		t.Fatal("expected at least one known path")
	}
	found := false
	for _, p := range paths {
		if p == "auth.accessTokenTTL" {
			found = true
		}
	}
	if !found {
		t.Error("auth.accessTokenTTL not in KnownRefPaths")
	}
}
