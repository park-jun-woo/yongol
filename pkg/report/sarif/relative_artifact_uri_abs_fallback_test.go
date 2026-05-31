//ff:func feature=report type=test control=sequence topic=sarif
//ff:what TestRelativeArtifactURI — empty/no-specsDir/rel-ok/abs-fallback/escape 분기 검증
package sarif

import (
	"path/filepath"
	"testing"
)

func TestRelativeArtifactURI_AbsFallback(t *testing.T) {
	dir := t.TempDir()
	absSpecs, _ := filepath.Abs(dir)
	file := filepath.Join(absSpecs, "sub", "x.ssac")

	got := relativeArtifactURI(file, "nonmatching-specs", absSpecs)
	if got != "sub/x.ssac" {
		t.Errorf("abs fallback: got %q, want sub/x.ssac", got)
	}
}
