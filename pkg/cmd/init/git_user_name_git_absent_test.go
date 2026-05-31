//ff:func feature=cli-init type=test control=sequence
//ff:what gitUserName — git 부재(에러경로) 및 호출 안정성 검증
package cliinit

import (
	"testing"
)

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
