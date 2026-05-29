//ff:func feature=funcspec type=test control=sequence
//ff:what collectPackageTypes — 존재하지 않는 경로는 SILENT-OK (진단 0, 결과 비어있음)

package funcspec

import (
	"path/filepath"
	"testing"
)

// TestCollectPackageTypesMissingDir — a non-existent path must be SILENT-OK (zero diagnostics).
func TestCollectPackageTypesMissingDir(t *testing.T) {
	dir := filepath.Join(t.TempDir(), "does-not-exist")
	result, diags := collectPackageTypes(dir)
	if len(diags) != 0 {
		t.Fatalf("missing dir should be SILENT-OK, got diags=%v", diags)
	}
	if len(result) != 0 {
		t.Errorf("result should be empty, got keys=%v", keysOf(result))
	}
}
