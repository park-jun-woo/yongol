//ff:func feature=orchestrator type=test control=sequence
//ff:what findYongolPkgRoot — YONGOL_SSAC_PKG 환경 override 우선 분기
package yongol

import (
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
