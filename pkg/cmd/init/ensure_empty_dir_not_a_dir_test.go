//ff:func feature=cli-init type=test control=sequence
//ff:what TestEnsureEmptyDir — 부재/파일/force/빈디렉/비어있지않음 분기 검증
package cliinit

import (
	"os"
	"path/filepath"
	"testing"
)

func TestEnsureEmptyDir_NotADir(t *testing.T) {
	f := filepath.Join(t.TempDir(), "file")
	if err := os.WriteFile(f, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := ensureEmptyDir(f, false); err == nil {
		t.Fatal("want error when path is a file")
	}
}
