//ff:func feature=external type=test control=sequence
//ff:what readSource — 존재하지 않는 파일 경로는 os.IsNotExist(err)

package external

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadSource_FileNotFound(t *testing.T) {
	path := filepath.Join(t.TempDir(), "does-not-exist.yaml")
	_, err := readSource(path)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !os.IsNotExist(err) {
		t.Errorf("expected os.IsNotExist(err) = true, got: %v", err)
	}
}
