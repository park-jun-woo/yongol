//ff:func feature=orchestrator type=test control=sequence
//ff:what TestTrySSaCPkgPath — sibling ssac/pkg 존재(경로 반환) / 부재 또는 파일("") 분기 검증
package yongol

import (
	"os"
	"path/filepath"
	"testing"
)

func TestTrySSaCPkgPath_Exists(t *testing.T) {
	base := t.TempDir()
	yongolRoot := filepath.Join(base, "yongol")
	ssacPkg := filepath.Join(base, "ssac", "pkg")
	if err := os.MkdirAll(yongolRoot, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(ssacPkg, 0o755); err != nil {
		t.Fatal(err)
	}
	got := trySSaCPkgPath(yongolRoot)
	if got != ssacPkg {
		t.Fatalf("got %q, want %q", got, ssacPkg)
	}
}
