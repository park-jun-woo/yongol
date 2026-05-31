//ff:func feature=gen-splitter type=test control=sequence
//ff:what zz_zerocov_test — splitter 패키지의 0% 커버리지 함수(cleanOriginal/preserveComments/isPreservedFile/SplitDirectory) 단위 테스트
package splitter

import (
	"path/filepath"
	"testing"
)

func TestSplitDirectory_NoDir_ZeroCov(t *testing.T) {
	// Non-existent directory → no-op, nil error.
	if err := SplitDirectory(filepath.Join(t.TempDir(), "absent"), ToolSQLC); err != nil {
		t.Fatalf("expected nil for absent dir, got %v", err)
	}
}
