//ff:func feature=orchestrator type=test control=sequence
//ff:what DetectSSOTs — 파일 경로를 루트로 주면 not a directory 에러
package yongol

import (
	"os"
	"path/filepath"
	"testing"
)

// TestDetectSSOTsNotADirectory asserts that pointing DetectSSOTs at a regular
// file returns an explicit "not a directory" error.
func TestDetectSSOTsNotADirectory(t *testing.T) {
	tmp := newTmpSpecsDir(t)
	f := filepath.Join(tmp, "not-a-dir.txt")
	if err := os.WriteFile(f, []byte("x"), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}

	_, err := DetectSSOTs(f)
	if err == nil {
		t.Fatalf("expected error when root is a file, got nil")
	}
}
