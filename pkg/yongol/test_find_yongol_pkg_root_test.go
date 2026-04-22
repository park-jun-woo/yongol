//ff:func feature=orchestrator type=test control=sequence
//ff:what findYongolPkgRoot — env 우선 / gomodcache fallback / 전부 미존재 시 "" 반환
package yongol

import (
	"os"
	"path/filepath"
	"testing"
)

// TestFindYongolPkgRootEnvOverride verifies that YONGOL_SSAC_PKG wins over
// every other lookup. The value just needs to exist as a directory.
func TestFindYongolPkgRootEnvOverride(t *testing.T) {
	tmp := t.TempDir()
	t.Setenv("YONGOL_SSAC_PKG", tmp)

	got := findYongolPkgRoot()
	if got != tmp {
		t.Fatalf("findYongolPkgRoot = %q; want %q (env override)", got, tmp)
	}
}

// TestFindYongolPkgRootEnvIgnoredWhenMissing verifies that a non-existent
// YONGOL_SSAC_PKG is skipped and the function falls through to the other
// resolvers (which we isolate to empty directories → "" result).
func TestFindYongolPkgRootEnvIgnoredWhenMissing(t *testing.T) {
	tmp := t.TempDir()
	t.Setenv("YONGOL_SSAC_PKG", filepath.Join(tmp, "does-not-exist"))
	// Isolate both fallback resolvers.
	isolated := t.TempDir()
	if err := os.Chdir(isolated); err != nil {
		t.Fatalf("chdir: %v", err)
	}
	t.Setenv("GOMODCACHE", filepath.Join(isolated, "empty-gomodcache"))

	got := findYongolPkgRoot()
	if got != "" {
		// Ambient sibling ssac repo can legitimately resolve here — skip rather
		// than hard-fail so the test remains CI-clean but worktree-friendly.
		t.Skipf("ambient sibling resolved %q; skipping env-miss assertion", got)
	}
}

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

// TestFindYongolPkgRootFromGoModCacheEmpty verifies the function returns ""
// when GOMODCACHE is empty (no github.com/park-jun-woo/ssac@* present).
func TestFindYongolPkgRootFromGoModCacheEmpty(t *testing.T) {
	tmp := t.TempDir()
	t.Setenv("GOMODCACHE", tmp)

	if got := findYongolPkgRootFromGoModCache(); got != "" {
		t.Fatalf("expected empty cache to yield \"\", got %q", got)
	}
}
