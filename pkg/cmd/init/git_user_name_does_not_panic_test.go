//ff:func feature=cli-init type=test control=sequence
//ff:what gitUserName — git 부재(에러경로) 및 호출 안정성 검증
package cliinit

import (
	"testing"
)

func TestGitUserNameDoesNotPanic(t *testing.T) {
	name, ok := gitUserName()
	if ok && name == "" {
		t.Error("ok=true but name is empty")
	}
	if !ok && name != "" {
		t.Errorf("ok=false but name=%q", name)
	}
}
