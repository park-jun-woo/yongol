//ff:func feature=cli-init type=test control=sequence
//ff:what TestCreateSkeletonDirs — 전체 디렉토리 생성 성공 / target이 파일이면 mkdir 에러 분기 검증
package cliinit

import (
	"testing"
)

func TestCreateSkeletonDirs_Success(t *testing.T) {
	target := t.TempDir()
	if err := createSkeletonDirs(target); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	assertSkeletonDirs(t, target)
}
