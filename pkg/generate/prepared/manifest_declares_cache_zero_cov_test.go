//ff:func feature=generate type=test control=sequence
//ff:what prepared 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package prepared

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestManifestDeclaresCache_ZeroCov(t *testing.T) {
	if manifestDeclaresCache(nil) {
		t.Error("nil should be false")
	}
	mc := &pmanifest.ProjectConfig{Cache: &pmanifest.BuiltinBackend{Backend: "redis"}}
	if !manifestDeclaresCache(bnFS(mc, nil)) {
		t.Error("declared cache should be true")
	}
}
