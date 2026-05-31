//ff:func feature=report type=test control=sequence topic=sarif
//ff:what TestRelativeArtifactURI — empty/no-specsDir/rel-ok/abs-fallback/escape 분기 검증
package sarif

import (
	"testing"
)

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
