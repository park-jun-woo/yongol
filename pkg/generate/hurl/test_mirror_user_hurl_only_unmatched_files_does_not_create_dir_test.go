//ff:func feature=gen-hurl type=test control=sequence
//ff:what TestMirrorUserHurlFiles_OnlyUnmatchedFilesDoesNotCreateDir — 매치 없으면 arts/tests/ 생성 안 함

package hurl

import (
	"os"
	"path/filepath"
	"testing"
)

// TestMirrorUserHurlFiles_OnlyUnmatchedFilesDoesNotCreateDir ensures
// tests/ dir with only ignored files (e.g. stray smoke.hurl) does not
// materialize an empty arts/tests/.
func TestMirrorUserHurlFiles_OnlyUnmatchedFilesDoesNotCreateDir(t *testing.T) {
	tmp := t.TempDir()
	specsDir := filepath.Join(tmp, "specs")
	artsDir := filepath.Join(tmp, "arts")
	testsSrc := filepath.Join(specsDir, "tests")
	if err := os.MkdirAll(testsSrc, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	writeFile(t, filepath.Join(testsSrc, "smoke.hurl"), "# stray\n")
	writeFile(t, filepath.Join(testsSrc, "notes.txt"), "noise\n")

	if err := mirrorUserHurlFiles(specsDir, artsDir); err != nil {
		t.Fatalf("mirrorUserHurlFiles: %v", err)
	}
	if _, err := os.Stat(filepath.Join(artsDir, "tests")); !os.IsNotExist(err) {
		t.Fatalf("expected arts/tests/ NOT to be created; stat err=%v", err)
	}
}
