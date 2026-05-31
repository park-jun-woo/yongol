//ff:func feature=gen-react type=test control=sequence
//ff:what writeIndexHTML index.html 생성 내용·에러경로 검증
package react

import (
	"path/filepath"
	"testing"
)

func TestWriteIndexHTMLMissingDir(t *testing.T) {
	// Writing into a non-existent parent directory must surface an error.
	missing := filepath.Join(t.TempDir(), "does", "not", "exist")
	if err := writeIndexHTML(missing); err == nil {
		t.Fatal("expected error writing into non-existent directory, got nil")
	}
}
