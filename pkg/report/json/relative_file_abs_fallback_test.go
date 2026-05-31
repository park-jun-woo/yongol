//ff:func feature=report type=test control=sequence topic=json
//ff:what TestRelativeFile — empty/no-specsDir/rel-ok/abs-fallback/escape 분기 검증
package json

import (
	"path/filepath"
	"testing"
)

func TestRelativeFile_AbsFallback(t *testing.T) {
	dir := t.TempDir()
	absSpecs, _ := filepath.Abs(dir)
	file := filepath.Join(absSpecs, "sub", "x.ssac")

	// specsDir is a *relative* token that won't match the absolute file via
	// plain Rel, so the abs fallback must kick in.
	got := relativeFile(file, "nonmatching-specs", absSpecs)
	if got != "sub/x.ssac" {
		t.Errorf("abs fallback: got %q, want sub/x.ssac", got)
	}
}
