//ff:func feature=cli-init type=test control=sequence
//ff:what TestEnsureEmptyDir — 부재/파일/force/빈디렉/비어있지않음 분기 검증
package cliinit

import (
	"testing"
)

func TestEnsureEmptyDir_EmptyOK(t *testing.T) {
	if err := ensureEmptyDir(t.TempDir(), false); err != nil {
		t.Fatalf("empty dir should be ok: %v", err)
	}
}
