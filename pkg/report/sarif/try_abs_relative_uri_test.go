//ff:func feature=report type=test control=sequence topic=sarif
//ff:what TestTryAbsRelativeURI — empty absSpecs / rebase 성공 / escape("..") 분기 검증
package sarif

import (
	"path/filepath"
	"testing"
)

func TestTryAbsRelativeURI(t *testing.T) {
	if got := tryAbsRelativeURI("anything", ""); got != "" {
		t.Errorf("empty absSpecs: got %q, want empty", got)
	}

	dir := t.TempDir()
	absSpecs, _ := filepath.Abs(dir)

	file := filepath.Join(absSpecs, "deep", "y.ssac")
	if got := tryAbsRelativeURI(file, absSpecs); got != "deep/y.ssac" {
		t.Errorf("rebase: got %q, want deep/y.ssac", got)
	}

	sibling := filepath.Join(filepath.Dir(absSpecs), "elsewhere", "z.ssac")
	if got := tryAbsRelativeURI(sibling, absSpecs); got != "" {
		t.Errorf("escape: got %q, want empty", got)
	}
}
