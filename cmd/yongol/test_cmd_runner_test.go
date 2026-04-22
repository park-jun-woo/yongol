//ff:func feature=cli type=test-helper control=sequence
//ff:what runCmd / zenflowSpecsDir — shared cobra CLI integration test helpers

package main

import (
	"bytes"
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

// runCmd executes the yongol cobra tree built by newRoot() with the given
// argv and returns captured stdout, stderr, and the Execute() error. Tests
// assert on the error to infer the intended exit code:
//
//   - err == nil                  → exit 0
//   - err is *usageError          → exit 2 (mirrors main.go errors.As branch)
//   - err != nil (other)          → exit 1
//
// Nothing in the cobra tree writes to os.Stdout/os.Stderr when SetOut/SetErr
// are wired — except versionCmd which uses fmt.Printf directly. The version
// test therefore captures via the os.Stdout redirect helper below.
func runCmd(t *testing.T, args ...string) (stdout, stderr string, err error) {
	t.Helper()
	root := newRoot()
	var sb, eb bytes.Buffer
	root.SetOut(&sb)
	root.SetErr(&eb)
	root.SetArgs(args)
	err = root.Execute()
	return sb.String(), eb.String(), err
}

// isUsageError reports whether err originated from usageArgs / the root
// FlagErrorFunc. Matches main.go's exit-code 2 branch.
func isUsageError(err error) bool {
	var ue *usageError
	return errors.As(err, &ue)
}

// repoRoot walks up from this test file to the yongol module root. The
// anchor is go.mod — every caller expects to see it at the top of the repo.
func repoRoot(t *testing.T) string {
	t.Helper()
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller(0) failed")
	}
	dir := filepath.Dir(thisFile)
	for i := 0; i < 10; i++ {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	t.Fatal("repo root (go.mod) not found by walking up from test file")
	return ""
}

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
