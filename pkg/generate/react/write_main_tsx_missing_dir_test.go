//ff:func feature=gen-react type=test control=sequence
//ff:what writeMainTSX main.tsx 생성 내용·에러경로 검증
package react

import (
	"path/filepath"
	"testing"
)

func TestWriteMainTSXMissingDir(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "no", "such", "dir")
	if err := writeMainTSX(missing); err == nil {
		t.Fatal("expected error writing into non-existent directory, got nil")
	}
}
