//ff:func feature=gen-react type=test control=sequence
//ff:what writeSessionStoreClaims — stores 디렉토리 생성 실패 시 에러 반환 검증
package react

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteSessionStoreClaims_MkdirError(t *testing.T) {
	// srcDir is a regular file, so MkdirAll(srcDir/stores) must fail
	// (a path component is not a directory) and surface the error.
	file := filepath.Join(t.TempDir(), "notadir")
	if err := os.WriteFile(file, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := writeSessionStoreClaims(file, "localStorage", true); err == nil {
		t.Fatal("expected error creating stores dir under a file, got nil")
	}
}
