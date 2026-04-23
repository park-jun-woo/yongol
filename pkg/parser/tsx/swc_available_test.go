//ff:func feature=tsx-parser type=test-helper control=sequence
//ff:what swcAvailable — @swc/core 설치 여부 확인 (없으면 t.Skip)

package tsx

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

// swcAvailable returns true when a @swc/core install is discoverable from
// YONGOL_SWC_PROJECT_DIR or the parent project. Tests skip when absent —
// CI and local dev without Node is expected to skip, not fail, parser tests.
func swcAvailable(t *testing.T) bool {
	t.Helper()
	if _, err := exec.LookPath("node"); err != nil {
		t.Skip("node not available on PATH; skipping tsx parser tests")
		return false
	}
	dir := os.Getenv("YONGOL_SWC_PROJECT_DIR")
	if dir == "" {
		t.Skip("YONGOL_SWC_PROJECT_DIR not set; skipping tsx parser tests (install @swc/core and point YONGOL_SWC_PROJECT_DIR to its parent)")
		return false
	}
	if _, err := os.Stat(filepath.Join(dir, "node_modules", "@swc", "core")); err != nil {
		t.Skipf("@swc/core not installed at %s; skipping", dir)
		return false
	}
	return true
}
