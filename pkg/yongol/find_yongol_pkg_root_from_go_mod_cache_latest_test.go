//ff:func feature=orchestrator type=test control=iteration dimension=1
//ff:what findYongolPkgRootFromGoModCache — semver 최신 ssac@<ver>/pkg 선택
package yongol

import (
	"os"
	"path/filepath"
	"testing"
)

// TestFindYongolPkgRootFromGoModCacheLatest asserts that the GOMODCACHE
// resolver picks the latest `ssac@<ver>` directory (semver descending) as
// long as it actually contains a `pkg/` subdirectory.
func TestFindYongolPkgRootFromGoModCacheLatest(t *testing.T) {
	tmp := t.TempDir()
	// Build: <tmp>/github.com/park-jun-woo/ssac@v0.1.0/pkg (older, skipped)
	//        <tmp>/github.com/park-jun-woo/ssac@v0.2.0/pkg (latest, selected)
	vendor := filepath.Join(tmp, "github.com", "park-jun-woo")
	for _, v := range []string{"ssac@v0.1.0", "ssac@v0.2.0"} {
		if err := os.MkdirAll(filepath.Join(vendor, v, "pkg"), 0o755); err != nil {
			t.Fatalf("mkdir %s: %v", v, err)
		}
	}

	t.Setenv("GOMODCACHE", tmp)
	got := findYongolPkgRootFromGoModCache()
	want := filepath.Join(vendor, "ssac@v0.2.0", "pkg")
	if got != want {
		t.Fatalf("findYongolPkgRootFromGoModCache = %q; want %q", got, want)
	}
}
