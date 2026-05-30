//ff:func feature=cli-init type=test control=sequence
//ff:what gitUserName — git 부재(에러경로) 및 호출 안정성 검증

package cliinit

import (
	"testing"
)

// TestGitUserNameGitAbsent isolates PATH so the `git` binary cannot be found.
// exec.Command then fails to start, and gitUserName must report (\"\", false)
// rather than panicking.
func TestGitUserNameGitAbsent(t *testing.T) {
	t.Setenv("PATH", t.TempDir())

	name, ok := gitUserName()
	if ok {
		t.Errorf("expected ok=false when git is unavailable, got name=%q", name)
	}
	if name != "" {
		t.Errorf("expected empty name when git is unavailable, got %q", name)
	}
}

// TestGitUserNameDoesNotPanic exercises the real environment's git config. The
// result is environment-dependent, so we only assert the invariant: when ok is
// true the name is non-empty, and when ok is false the name is empty.
func TestGitUserNameDoesNotPanic(t *testing.T) {
	name, ok := gitUserName()
	if ok && name == "" {
		t.Error("ok=true but name is empty")
	}
	if !ok && name != "" {
		t.Errorf("ok=false but name=%q", name)
	}
}
