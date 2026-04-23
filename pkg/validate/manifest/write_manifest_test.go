//ff:func feature=validate type=test-helper control=sequence topic=manifest-auth
//ff:what writeManifest — tmp dir 에 manifest.yaml 기록 후 경로 반환

package manifest

import (
	"os"
	"path/filepath"
	"testing"
)

// writeManifest writes manifest.yaml into a tmp dir and returns the dir.
func writeManifest(t *testing.T, body string) string {
	t.Helper()
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "manifest.yaml"), []byte(body), 0o644); err != nil {
		t.Fatalf("write manifest: %v", err)
	}
	return dir
}
