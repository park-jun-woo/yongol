//ff:func feature=report type=test control=sequence topic=json
//ff:what TestRelativeFile — empty/no-specsDir/rel-ok/abs-fallback/escape 분기 검증
package json

import (
	"testing"
)

func TestRelativeFile(t *testing.T) {
	// empty file → empty string.
	if got := relativeFile("", "specs", "/abs/specs"); got != "" {
		t.Errorf("empty file: got %q, want empty", got)
	}

	// empty specsDir → ToSlash(file) only.
	if got := relativeFile("a/b/c.ssac", "", ""); got != "a/b/c.ssac" {
		t.Errorf("empty specsDir: got %q, want a/b/c.ssac", got)
	}

	// Relative-rebase branch: file under specsDir resolves to a clean rel path.
	if got := relativeFile("specs/auth/login.ssac", "specs", "/abs/specs"); got != "auth/login.ssac" {
		t.Errorf("rel-ok: got %q, want auth/login.ssac", got)
	}

	// Escape branch ("..") falls through: with no usable abs rebase the raw
	// slashed path is returned.
	got := relativeFile("other/x.ssac", "specs", "")
	if got != "other/x.ssac" {
		t.Errorf("escape fallback: got %q, want other/x.ssac", got)
	}
}
