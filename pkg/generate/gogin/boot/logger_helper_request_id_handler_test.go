//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestWriteRequestIDHandlerFile — request_id handler 파일 생성 + skip/error 경로 검증

package boot

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWriteRequestIDHandlerFile(t *testing.T) {
	t.Run("EmptyModuleSkips", func(t *testing.T) {
		dir := t.TempDir()
		if err := writeRequestIDHandlerFile(dir, ""); err != nil {
			t.Fatalf("expected nil for empty module, got: %v", err)
		}
		if _, err := os.Stat(filepath.Join(dir, "backend")); !os.IsNotExist(err) {
			t.Errorf("expected no output for empty module, stat err: %v", err)
		}
	})

	t.Run("WritesAllHandlerFiles", func(t *testing.T) {
		dir := t.TempDir()
		if err := writeRequestIDHandlerFile(dir, "example.com/app"); err != nil {
			t.Fatalf("writeRequestIDHandlerFile error: %v", err)
		}
		cmdDir := filepath.Join(dir, "backend", "cmd")
		for _, name := range []string{
			"request_id_handler.go",
			"request_id_handler_handle.go",
			"request_id_handler_with_attrs.go",
			"request_id_handler_with_group.go",
		} {
			info, err := os.Stat(filepath.Join(cmdDir, name))
			if err != nil {
				t.Errorf("expected %s: %v", name, err)
				continue
			}
			if info.Size() == 0 {
				t.Errorf("%s is empty", name)
			}
		}
	})

	t.Run("MkdirError", func(t *testing.T) {
		dir := t.TempDir()
		// backend/cmd parent (backend) is a regular file -> MkdirAll fails.
		if err := os.WriteFile(filepath.Join(dir, "backend"), []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := writeRequestIDHandlerFile(dir, "example.com/app"); err == nil {
			t.Errorf("expected MkdirAll error, got nil")
		}
	})

	t.Run("WriteFileError", func(t *testing.T) {
		dir := t.TempDir()
		cmdDir := filepath.Join(dir, "backend", "cmd")
		if err := os.MkdirAll(cmdDir, 0o755); err != nil {
			t.Fatalf("setup: %v", err)
		}
		// Pre-create every target as a directory so WriteFile fails for whichever is hit first.
		for _, name := range []string{
			"request_id_handler.go",
			"request_id_handler_handle.go",
			"request_id_handler_with_attrs.go",
			"request_id_handler_with_group.go",
		} {
			if err := os.MkdirAll(filepath.Join(cmdDir, name), 0o755); err != nil {
				t.Fatalf("setup: %v", err)
			}
		}
		err := writeRequestIDHandlerFile(dir, "example.com/app")
		if err == nil || !strings.Contains(err.Error(), "write ") {
			t.Errorf("expected write error, got: %v", err)
		}
	})
}
