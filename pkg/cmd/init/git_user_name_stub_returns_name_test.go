//ff:func feature=cli-init type=test control=sequence
//ff:what TestGitUserNameStub — 가짜 git 바이너리로 이름 반환 성공 / 빈 출력 false 분기 검증
package cliinit

import (
	"testing"
)

func TestGitUserNameStub_ReturnsName(t *testing.T) {
	writeFakeGit(t, "Alice Example\n")
	name, ok := gitUserName()
	if !ok || name != "Alice Example" {
		t.Errorf("got (%q, %v), want (\"Alice Example\", true)", name, ok)
	}
}
