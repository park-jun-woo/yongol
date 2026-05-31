//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestRemoveStaleGoModFiles — 기존 파일 제거 / 부재 무시 / 제거 실패 에러 분기 검증
package gogin

import (
	"testing"
)

func TestRemoveStaleGoModFiles_NoFiles(t *testing.T) {
	// Empty dir -> os.Remove returns IsNotExist which is tolerated.
	if err := removeStaleGoModFiles(t.TempDir()); err != nil {
		t.Errorf("expected nil when no stale files, got: %v", err)
	}
}
