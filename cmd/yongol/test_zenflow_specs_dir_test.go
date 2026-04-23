//ff:func feature=cli type=test-helper control=sequence
//ff:what zenflowSpecsDir — dummys/zenflow/try-02/specs 경로 반환 (없으면 skip)

package main

import (
	"os"
	"path/filepath"
	"testing"
)

// zenflowSpecsDir returns the absolute path to dummys/zenflow/try-02/specs or
// calls t.Skip when the dummy tree is absent (CI trimmed / non-dev checkout).
func zenflowSpecsDir(t *testing.T) string {
	t.Helper()
	dir := filepath.Join(repoRoot(t), "dummys", "zenflow", "try-02", "specs")
	info, err := os.Stat(dir)
	if err != nil || !info.IsDir() {
		t.Skipf("zenflow dummy specs not available at %s", dir)
	}
	return dir
}
