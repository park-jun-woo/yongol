//ff:func feature=gen-splitter type=test control=sequence
//ff:what zz_zerocov_test — splitter 패키지의 0% 커버리지 함수(cleanOriginal/preserveComments/isPreservedFile/SplitDirectory) 단위 테스트
package splitter

import (
	"path/filepath"
	"testing"
)

func TestCleanOriginal_ReadDirError_ZeroCov(t *testing.T) {
	if err := cleanOriginal(filepath.Join(t.TempDir(), "does-not-exist"), ToolSQLC, nil); err == nil {
		t.Fatal("expected error for missing dir")
	}
}
