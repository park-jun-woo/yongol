//ff:func feature=report type=test control=selection topic=sarif
//ff:what TestRelativeArtifactURI — empty/no-specsDir/rel-ok/abs-fallback/escape 분기 검증
package sarif

import (
	"path/filepath"
	"testing"
)

// TestRelativeArtifactURI covers each branch of relativeArtifactURI.
func TestRelativeArtifactURI(t *testing.T) {
	if got := relativeArtifactURI("", "specs", "/abs/specs"); got != "" {
		t.Errorf("empty file: got %q, want empty", got)
	}
	if got := relativeArtifactURI("a/b/c.ssac", "", ""); got != "a/b/c.ssac" {
		t.Errorf("empty specsDir: got %q, want a/b/c.ssac", got)
	}
	if got := relativeArtifactURI("specs/auth/login.ssac", "specs", "/abs/specs"); got != "auth/login.ssac" {
		t.Errorf("rel-ok: got %q, want auth/login.ssac", got)
	}
	// Escape + no usable abs root → raw slashed path.
	if got := relativeArtifactURI("other/x.ssac", "specs", ""); got != "other/x.ssac" {
		t.Errorf("escape fallback: got %q, want other/x.ssac", got)
	}
}

// TestRelativeArtifactURI_AbsFallback exercises the tryAbsRelativeURI success
// path where plain Rel fails but the absolute rebase succeeds.
func TestRelativeArtifactURI_AbsFallback(t *testing.T) {
	dir := t.TempDir()
	absSpecs, _ := filepath.Abs(dir)
	file := filepath.Join(absSpecs, "sub", "x.ssac")

	got := relativeArtifactURI(file, "nonmatching-specs", absSpecs)
	if got != "sub/x.ssac" {
		t.Errorf("abs fallback: got %q, want sub/x.ssac", got)
	}
}
