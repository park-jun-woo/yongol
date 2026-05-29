//ff:func feature=report type=test control=selection topic=json
//ff:what TestTryAbsRelativeFile — empty absSpecs / rebase 성공 / escape("..") 분기 검증
package json

import (
	"path/filepath"
	"testing"
)

// TestTryAbsRelativeFile covers the three branches: empty absSpecs guard,
// successful rebase, and the ".." escape rejection.
func TestTryAbsRelativeFile(t *testing.T) {
	// empty absSpecs → "".
	if got := tryAbsRelativeFile("anything", ""); got != "" {
		t.Errorf("empty absSpecs: got %q, want empty", got)
	}

	dir := t.TempDir()
	absSpecs, _ := filepath.Abs(dir)

	// Success: file inside absSpecs.
	file := filepath.Join(absSpecs, "deep", "y.ssac")
	if got := tryAbsRelativeFile(file, absSpecs); got != "deep/y.ssac" {
		t.Errorf("rebase: got %q, want deep/y.ssac", got)
	}

	// Escape: a sibling directory rebases to a "../..." path → rejected ("").
	sibling := filepath.Join(filepath.Dir(absSpecs), "elsewhere", "z.ssac")
	if got := tryAbsRelativeFile(sibling, absSpecs); got != "" {
		t.Errorf("escape: got %q, want empty", got)
	}
}
