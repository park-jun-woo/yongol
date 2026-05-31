//ff:func feature=funcspec type=test control=sequence
//ff:what funcspec 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package funcspec

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestLoadTypeMapForDir_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "m.go"), []byte(bnStructSrc), 0644); err != nil {
		t.Fatal(err)
	}
	cache := map[string]map[string][]Field{}
	seen := map[string]struct{}{}
	tm, diags := loadTypeMapForDir(dir, cache, seen, nil)
	if tm == nil {
		t.Errorf("expected type map")
	}
	_ = diags
	// second call hits cache.
	tm2, _ := loadTypeMapForDir(dir, cache, seen, []diagnostic.Diagnostic{})
	if tm2 == nil {
		t.Errorf("expected cached type map")
	}
}
