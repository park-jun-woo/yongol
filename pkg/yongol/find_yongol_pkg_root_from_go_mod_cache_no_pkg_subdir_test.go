//ff:func feature=orchestrator type=test control=sequence
//ff:what findYongolPkgRootFromGoModCache — 후보 없음 / pkg 없는 ssac 디렉토리 분기 검증
package yongol

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindYongolPkgRootFromGoModCache_NoPkgSubdir(t *testing.T) {
	tmp := t.TempDir()
	vendor := filepath.Join(tmp, "github.com", "park-jun-woo")
	if err := os.MkdirAll(filepath.Join(vendor, "ssac@v0.9.0"), 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("GOMODCACHE", tmp)
	if got := findYongolPkgRootFromGoModCache(); got != "" {
		t.Fatalf("expected \"\" when ssac dir has no pkg/, got %q", got)
	}
}
