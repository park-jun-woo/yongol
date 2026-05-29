//ff:func feature=orchestrator type=test control=sequence
//ff:what DetectSSOTs — 존재하지 않는 경로는 os.Stat 에러 전파
package yongol

import (
	"path/filepath"
	"testing"
)

// TestDetectSSOTsNonExistentRoot asserts that a non-existent path bubbles up
// an os.Stat error rather than crashing.
func TestDetectSSOTsNonExistentRoot(t *testing.T) {
	tmp := newTmpSpecsDir(t)
	_, err := DetectSSOTs(filepath.Join(tmp, "does-not-exist"))
	if err == nil {
		t.Fatalf("expected error when root missing, got nil")
	}
}
