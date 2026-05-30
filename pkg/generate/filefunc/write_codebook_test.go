//ff:func feature=gen-filefunc type=test control=sequence
//ff:what TestWriteCodebook — Codebook YAML 직렬화/기록 + mkdir 에러 경로 검증

package filefunc

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestWriteCodebook(t *testing.T) {
	t.Run("WritesYAMLWithHeader", func(t *testing.T) {
		path := filepath.Join(t.TempDir(), "nested", "codebook.yaml")
		book := Codebook{}
		if err := writeCodebook(path, book); err != nil {
			t.Fatalf("writeCodebook error: %v", err)
		}
		got, err := os.ReadFile(path)
		if err != nil {
			t.Fatalf("read codebook: %v", err)
		}
		if !bytes.HasPrefix(got, []byte(codebookHeader)) {
			t.Errorf("expected header prefix, got:\n%s", got)
		}
		if !bytes.Contains(got, []byte("required:")) {
			t.Errorf("expected serialized 'required:' key, got:\n%s", got)
		}
	})

	t.Run("WriteFileFails", func(t *testing.T) {
		// Target path is a directory -> os.WriteFile fails.
		dir := t.TempDir()
		target := filepath.Join(dir, "codebook.yaml")
		if err := os.MkdirAll(target, 0o755); err != nil {
			t.Fatalf("setup: %v", err)
		}
		err := writeCodebook(target, Codebook{})
		if err == nil {
			t.Errorf("expected write codebook error, got nil")
		}
	})

	t.Run("MkdirFails", func(t *testing.T) {
		// Parent path component is a regular file -> MkdirAll fails.
		base := t.TempDir()
		fp := filepath.Join(base, "file")
		if err := os.WriteFile(fp, []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		err := writeCodebook(filepath.Join(fp, "sub", "codebook.yaml"), Codebook{})
		if err == nil {
			t.Errorf("expected mkdir error, got nil")
		}
	})
}
