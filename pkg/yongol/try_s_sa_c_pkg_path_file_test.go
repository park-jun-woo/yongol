//ff:func feature=orchestrator type=test control=sequence
//ff:what TestTrySSaCPkgPath — sibling ssac/pkg 존재(경로 반환) / 부재 또는 파일("") 분기 검증
package yongol

import (
	"os"
	"path/filepath"
	"testing"
)

func TestTrySSaCPkgPath_File(t *testing.T) {
	base := t.TempDir()
	yongolRoot := filepath.Join(base, "yongol")
	if err := os.MkdirAll(filepath.Join(base, "ssac"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(yongolRoot, 0o755); err != nil {
		t.Fatal(err)
	}
	// ssac/pkg exists but is a regular file, not a directory → "".
	if err := os.WriteFile(filepath.Join(base, "ssac", "pkg"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if got := trySSaCPkgPath(yongolRoot); got != "" {
		t.Fatalf("expected \"\" when sibling is a file, got %q", got)
	}
}
