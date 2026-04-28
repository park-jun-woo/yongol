//ff:func feature=gen-hurl type=test-helper control=sequence
//ff:what mustWrite — 테스트용 파일 기록 헬퍼

package hurl_mirror

import (
	"os"
	"path/filepath"
	"testing"
)

// mustWrite creates parent directories and writes content to path. Used by
// MirrorSpecsTests fixture helpers; fails the test on any filesystem error.
func mustWrite(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("mkdir %s: %v", filepath.Dir(path), err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write %s: %v", path, err)
	}
}
