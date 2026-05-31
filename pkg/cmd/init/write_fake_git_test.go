//ff:func feature=cli-init type=test control=sequence
//ff:what TestGitUserNameStub — 가짜 git 바이너리로 이름 반환 성공 / 빈 출력 false 분기 검증
package cliinit

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func writeFakeGit(t *testing.T, output string) {
	t.Helper()
	if runtime.GOOS == "windows" {
		t.Skip("fake git shim is POSIX-only")
	}
	dir := t.TempDir()
	script := "#!/bin/sh\nprintf '%s' \"" + output + "\"\n"
	path := filepath.Join(dir, "git")
	if err := os.WriteFile(path, []byte(script), 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("PATH", dir)
}
