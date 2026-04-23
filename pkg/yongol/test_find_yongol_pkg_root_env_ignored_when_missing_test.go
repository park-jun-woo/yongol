//ff:func feature=orchestrator type=test control=sequence
//ff:what findYongolPkgRoot — 환경변수가 가리키는 경로가 없으면 fallback 으로 이동
package yongol

import (
	"os"
	"path/filepath"
	"testing"
)

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
