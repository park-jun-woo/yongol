//ff:func feature=cli-init type=test control=sequence
//ff:what TestEnsureEmptyDir — 부재/파일/force/빈디렉/비어있지않음 분기 검증
package cliinit

import (
	"path/filepath"
	"testing"
)

func TestEnsureEmptyDir_NotExist(t *testing.T) {
	if err := ensureEmptyDir(filepath.Join(t.TempDir(), "nope"), false); err != nil {
		t.Fatalf("missing dir should be ok: %v", err)
	}
}
