//ff:func feature=cli-init type=test control=sequence
//ff:what TestEnsureEmptyDir — 부재/파일/force/빈디렉/비어있지않음 분기 검증
package cliinit

import (
	"os"
	"path/filepath"
	"testing"
)

func TestEnsureEmptyDir_ForceOverridesNonEmpty(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "f"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := ensureEmptyDir(dir, true); err != nil {
		t.Fatalf("force should override: %v", err)
	}
}
