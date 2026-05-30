//ff:func feature=gen-filefunc type=test control=sequence
//ff:what TestMergeInternalDirs — internal 하위 디렉토리 병합 + 미존재 무시 검증

package filefunc

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMergeInternalDirs(t *testing.T) {
	t.Run("MissingDirNoop", func(t *testing.T) {
		dst := map[string]string{}
		mergeInternalDirs(dst, filepath.Join(t.TempDir(), "does-not-exist"))
		if len(dst) != 0 {
			t.Errorf("expected no-op for missing dir, got: %v", dst)
		}
	})

	t.Run("MergesSubdirs", func(t *testing.T) {
		internalDir := t.TempDir()
		if err := os.MkdirAll(filepath.Join(internalDir, "auth"), 0o755); err != nil {
			t.Fatalf("setup: %v", err)
		}
		if err := os.WriteFile(filepath.Join(internalDir, "x.go"), []byte("x"), 0o644); err != nil {
			t.Fatalf("setup: %v", err)
		}
		dst := map[string]string{}
		mergeInternalDirs(dst, internalDir)
		if _, ok := dst["auth"]; !ok {
			t.Errorf("expected auth subdir merged: %v", dst)
		}
		if _, ok := dst["x.go"]; ok {
			t.Errorf("file should not be merged")
		}
	})
}
