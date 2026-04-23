//ff:func feature=orchestrator type=test control=sequence
//ff:what findYongolPkgRootFromGoModCache — 빈 GOMODCACHE 는 "" 반환
package yongol

import (
	"testing"
)

// TestFindYongolPkgRootFromGoModCacheEmpty verifies the function returns ""
// when GOMODCACHE is empty (no github.com/park-jun-woo/ssac@* present).
func TestFindYongolPkgRootFromGoModCacheEmpty(t *testing.T) {
	tmp := t.TempDir()
	t.Setenv("GOMODCACHE", tmp)

	if got := findYongolPkgRootFromGoModCache(); got != "" {
		t.Fatalf("expected empty cache to yield \"\", got %q", got)
	}
}
