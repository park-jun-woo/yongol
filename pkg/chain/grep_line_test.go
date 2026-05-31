//ff:func feature=chain type=test control=sequence
//ff:what grepLine 가 substr 포함 첫 줄 번호 / 미발견 / 파일없음을 올바르게 반환하는지 검증
package chain

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGrepLine(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "sample.txt")
	content := "first line\nsecond NEEDLE here\nthird line\n"
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}

	if got := grepLine(path, "NEEDLE"); got != 2 {
		t.Errorf("found substr: got line %d, want 2", got)
	}
	if got := grepLine(path, "first"); got != 1 {
		t.Errorf("first line: got %d, want 1", got)
	}
	if got := grepLine(path, "missing"); got != 0 {
		t.Errorf("missing substr: got %d, want 0", got)
	}
	if got := grepLine(filepath.Join(dir, "nope.txt"), "x"); got != 0 {
		t.Errorf("missing file: got %d, want 0", got)
	}
}
