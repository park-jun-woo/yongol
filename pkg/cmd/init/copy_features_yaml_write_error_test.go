//ff:func feature=cli-init type=test control=sequence
//ff:what TestCopyFeaturesYAML — 정상 복사 / src 읽기 에러 / dest 쓰기(specs 부재) 에러 분기 검증
package cliinit

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCopyFeaturesYAML_WriteError(t *testing.T) {
	// targetDir has no specs/ subdir -> dest write fails.
	src := filepath.Join(t.TempDir(), "features.yaml")
	if err := os.WriteFile(src, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := copyFeaturesYAML(t.TempDir(), src); err == nil {
		t.Fatal("want write error when specs/ missing")
	}
}
