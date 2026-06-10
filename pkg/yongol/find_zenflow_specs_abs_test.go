//ff:func feature=orchestrator type=test-helper control=sequence
//ff:what findZenflowSpecsAbs — 공용 zenflow fixture 절대경로 해석
package yongol

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

// findZenflowSpecsAbs resolves the absolute path of the shared zenflow fixture
// (`examples/zenflow/try-02/specs`) relative to *this* test file, so the tests
// work regardless of the caller's CWD and survive repo relocation.
func findZenflowSpecsAbs(t *testing.T) string {
	t.Helper()
	_, thisFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatalf("runtime.Caller(0) failed")
	}
	// thisFile: <repo>/pkg/yongol/test_find_zenflow_specs_abs_test.go
	pkgYongol := filepath.Dir(thisFile)
	repoRoot := filepath.Dir(filepath.Dir(pkgYongol))
	specs := filepath.Join(repoRoot, "examples", "zenflow", "try-02", "specs")
	if _, err := os.Stat(specs); err != nil {
		t.Skipf("zenflow fixture unavailable at %s: %v", specs, err)
	}
	return specs
}
