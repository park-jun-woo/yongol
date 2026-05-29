//ff:func feature=cli type=test-helper control=sequence
//ff:what opus47SpecsDir — opus4_7 example specs 디렉토리 경로 반환 (없으면 skip)

package main

import (
	"os"
	"path/filepath"
	"testing"
)

// opus47SpecsDir returns the opus4_7 example specs directory, skipping if absent.
func opus47SpecsDir(t *testing.T) string {
	t.Helper()
	dir := filepath.Join(repoRoot(t), "examples", "zenflow", "opus4_7", "specs")
	if _, err := os.Stat(dir); err != nil {
		t.Skipf("opus4_7 specs not available: %v", err)
	}
	return dir
}
