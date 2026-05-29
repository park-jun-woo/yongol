//ff:func feature=report type=test control=selection topic=json
//ff:what TestRelativeFile — empty/no-specsDir/rel-ok/abs-fallback/escape 분기 검증
package json

import (
	"path/filepath"
	"testing"
)

// TestRelativeFile covers each branch of relativeFile.
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

// TestRelativeFile_AbsFallback exercises the tryAbsRelativeFile success path
// where the plain Rel fails/escapes but an absolute-vs-absolute rebase wins.
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
