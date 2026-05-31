//ff:func feature=gen-splitter type=test control=sequence
//ff:what zz_zerocov_test — splitter 패키지의 0% 커버리지 함수(cleanOriginal/preserveComments/isPreservedFile/SplitDirectory) 단위 테스트
package splitter

import (
	"testing"
)

func TestSplitDirectory_SplitFileError_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	// A models.go that fails to parse → SplitFile returns an error which
	// SplitDirectory wraps and returns.
	writeFileZeroCov(t, dir, "models.go", "package db\nthis is not valid go\n")
	if err := SplitDirectory(dir, ToolSQLC); err == nil {
		t.Fatal("expected split error to propagate")
	}
}
