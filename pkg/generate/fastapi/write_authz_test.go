//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestWriteAuthz — FastAPI authz stub 파일 기록 검증

package fastapi

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestWriteAuthz(t *testing.T) {
	t.Run("WritesAuthzStub", func(t *testing.T) {
		appDir := t.TempDir()
		if err := writeAuthz(appDir); err != nil {
			t.Fatalf("writeAuthz error: %v", err)
		}
		data, err := os.ReadFile(filepath.Join(appDir, "dependencies", "authz.py"))
		if err != nil {
			t.Fatalf("expected authz.py: %v", err)
		}
		if !strings.Contains(string(data), "async def authz_check") {
			t.Errorf("authz.py missing authz_check definition")
		}
	})

	t.Run("MkdirFails", func(t *testing.T) {
		base := t.TempDir()
		filePath := filepath.Join(base, "file")
		if err := os.WriteFile(filePath, []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		// appDir is a regular file -> MkdirAll(depsDir) fails.
		if err := writeAuthz(filePath); err == nil {
			t.Errorf("expected MkdirAll error, got nil")
		}
	})
}
