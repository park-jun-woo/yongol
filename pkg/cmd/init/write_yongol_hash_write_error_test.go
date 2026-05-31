//ff:func feature=cli-init type=test control=sequence
//ff:what TestWriteYongolHash — 정상 해시 기록 / features 읽기 에러 / specs 부재 write 에러 분기 검증
package cliinit

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteYongolHash_WriteError(t *testing.T) {
	feat := filepath.Join(t.TempDir(), "features.yaml")
	if err := os.WriteFile(feat, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	// target has no specs/ subdir -> write fails.
	if err := writeYongolHash(t.TempDir(), feat); err == nil {
		t.Fatal("want write error when specs/ missing")
	}
}
