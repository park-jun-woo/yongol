//ff:func feature=external type=test control=sequence
//ff:what readSource — 로컬 파일 경로는 파일 내용 반환

package external

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadSource_FilePath(t *testing.T) {
	tmp := t.TempDir()
	path := filepath.Join(tmp, "spec.yaml")
	content := []byte("openapi: 3.0.0\n")
	if err := os.WriteFile(path, content, 0644); err != nil {
		t.Fatalf("setup write: %v", err)
	}

	data, err := readSource(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if string(data) != string(content) {
		t.Errorf("content mismatch: got %q want %q", string(data), string(content))
	}
}
