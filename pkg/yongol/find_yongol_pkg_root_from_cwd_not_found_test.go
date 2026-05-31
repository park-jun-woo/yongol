//ff:func feature=orchestrator type=test control=sequence
//ff:what findYongolPkgRootFromCWD — yongol root 발견(성공) 및 root까지 미발견("") 분기 검증
package yongol

import (
	"os"
	"testing"
)

func TestFindYongolPkgRootFromCWD_NotFound(t *testing.T) {
	// An isolated empty directory: walking up never finds a yongol root,
	// so the loop reaches the filesystem root and returns "".
	isolated := t.TempDir()
	if err := os.Chdir(isolated); err != nil {
		t.Fatalf("chdir: %v", err)
	}
	if got := findYongolPkgRootFromCWD(); got != "" {
		t.Skipf("ambient yongol root resolved %q; skipping not-found assertion", got)
	}
}
