//ff:func feature=cli-init type=test control=sequence
//ff:what TestWriteSkeletonFiles — 디렉토리 준비 후 전체 파일 기록 성공 / 디렉 부재 시 에러 분기 검증
package cliinit

import (
	"testing"
)

func TestWriteSkeletonFiles_Error(t *testing.T) {
	// Dirs not created -> first nested write fails (e.g. specs/api/openapi.yaml).
	target := t.TempDir()
	if err := writeSkeletonFiles(target, templateData{}); err == nil {
		t.Fatal("want error when skeleton dirs are missing")
	}
}
