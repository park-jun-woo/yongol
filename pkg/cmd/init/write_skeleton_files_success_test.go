//ff:func feature=cli-init type=test control=sequence
//ff:what TestWriteSkeletonFiles — 디렉토리 준비 후 전체 파일 기록 성공 / 디렉 부재 시 에러 분기 검증
package cliinit

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteSkeletonFiles_Success(t *testing.T) {
	target := t.TempDir()
	if err := createSkeletonDirs(target); err != nil {
		t.Fatal(err)
	}
	data := templateData{ProjectID: "App", ProjectIDNormalized: "app", Description: "d", Module: "github.com/x/app"}
	if err := writeSkeletonFiles(target, data); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// manifest.yaml is one of the rendered skeleton files.
	if _, err := os.Stat(filepath.Join(target, "specs", "manifest.yaml")); err != nil {
		t.Errorf("manifest not written: %v", err)
	}
}
