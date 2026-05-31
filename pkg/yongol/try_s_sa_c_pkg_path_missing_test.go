//ff:func feature=orchestrator type=test control=sequence
//ff:what TestTrySSaCPkgPath — sibling ssac/pkg 존재(경로 반환) / 부재 또는 파일("") 분기 검증
package yongol

import (
	"os"
	"path/filepath"
	"testing"
)

func TestTrySSaCPkgPath_Missing(t *testing.T) {
	base := t.TempDir()
	yongolRoot := filepath.Join(base, "yongol")
	if err := os.MkdirAll(yongolRoot, 0o755); err != nil {
		t.Fatal(err)
	}
	// No sibling ssac/pkg → "".
	if got := trySSaCPkgPath(yongolRoot); got != "" {
		t.Fatalf("expected \"\" when sibling missing, got %q", got)
	}
}
