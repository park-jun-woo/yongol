//ff:func feature=generate type=test control=sequence
//ff:what prepared 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package prepared

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestManifestDeclaresSession_ZeroCov(t *testing.T) {
	if manifestDeclaresSession(nil) {
		t.Error("nil should be false")
	}
	mc := &pmanifest.ProjectConfig{Session: &pmanifest.BuiltinBackend{Backend: "memory"}}
	if !manifestDeclaresSession(bnFS(mc, nil)) {
		t.Error("declared session should be true")
	}
}
