//ff:func feature=cli-init type=test control=sequence
//ff:what TestCreateSkeletonDirs — 전체 디렉토리 생성 성공 / target이 파일이면 mkdir 에러 분기 검증
package cliinit

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCreateSkeletonDirs_MkdirError(t *testing.T) {
	// Make targetDir a regular file so MkdirAll under it fails.
	parent := t.TempDir()
	target := filepath.Join(parent, "asfile")
	if err := os.WriteFile(target, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := createSkeletonDirs(target); err == nil {
		t.Fatal("want mkdir error when target is a file")
	}
}
