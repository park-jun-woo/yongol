//ff:func feature=generate type=test control=sequence
//ff:what prepared 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package prepared

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestCacheBackendFor_ZeroCov(t *testing.T) {
	mc := &pmanifest.ProjectConfig{Cache: &pmanifest.BuiltinBackend{Backend: "redis"}}
	if c := cacheBackendFor(bnFS(mc, nil)); c == nil || c.Backend != "redis" {
		t.Errorf("declared: %#v", c)
	}
	if c := cacheBackendFor(bnFS(nil, []ssac.ServiceFunc{bnCallFunc("cache.")})); c == nil || c.Backend != "memory" {
		t.Errorf("ssac default: %#v", c)
	}
	if c := cacheBackendFor(bnFS(nil, nil)); c != nil {
		t.Errorf("unused should be nil")
	}
}
