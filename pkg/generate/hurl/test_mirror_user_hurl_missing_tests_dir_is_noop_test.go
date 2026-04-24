//ff:func feature=gen-hurl type=test control=sequence
//ff:what TestMirrorUserHurlFiles_MissingTestsDirIsNoop — specs/tests/ 없으면 no-op

package hurl

import (
	"os"
	"path/filepath"
	"testing"
)

// TestMirrorUserHurlFiles_MissingTestsDirIsNoop ensures projects
// without specs/tests/ still succeed without creating arts/tests/.
func TestMirrorUserHurlFiles_MissingTestsDirIsNoop(t *testing.T) {
	tmp := t.TempDir()
	specsDir := filepath.Join(tmp, "specs")
	artsDir := filepath.Join(tmp, "arts")
	if err := os.MkdirAll(specsDir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	if err := mirrorUserHurlFiles(specsDir, artsDir); err != nil {
		t.Fatalf("mirrorUserHurlFiles: %v", err)
	}
	if _, err := os.Stat(filepath.Join(artsDir, "tests")); !os.IsNotExist(err) {
		t.Fatalf("expected no arts/tests/ when specs/tests/ missing; stat err=%v", err)
	}
}
