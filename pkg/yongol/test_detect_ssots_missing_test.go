//ff:func feature=orchestrator type=test control=sequence
//ff:what DetectSSOTs — 루트 자체가 없거나 파일이면 에러, 빈 디렉토리면 빈 슬라이스
package yongol

import (
	"os"
	"path/filepath"
	"testing"
)

// TestDetectSSOTsEmptyDir asserts that a completely empty specs root returns
// no detected SSOTs and no error. All kinds are SSOTAbsent and therefore
// omitted by design.
func TestDetectSSOTsEmptyDir(t *testing.T) {
	tmp := newTmpSpecsDir(t)

	detected, err := DetectSSOTs(tmp)
	if err != nil {
		t.Fatalf("DetectSSOTs: %v", err)
	}
	if len(detected) != 0 {
		t.Fatalf("expected 0 detected on empty dir, got %d: %+v", len(detected), detected)
	}
}

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

// TestDetectSSOTsNonExistentRoot asserts that a non-existent path bubbles up
// an os.Stat error rather than crashing.
func TestDetectSSOTsNonExistentRoot(t *testing.T) {
	tmp := newTmpSpecsDir(t)
	_, err := DetectSSOTs(filepath.Join(tmp, "does-not-exist"))
	if err == nil {
		t.Fatalf("expected error when root missing, got nil")
	}
}
