//ff:func feature=funcspec type=test control=sequence
//ff:what fillMissingFields — 같은 dir 중복 호출 시 진단은 1회만 방출 (seenDir dedup)

package funcspec

import (
	"os"
	"path/filepath"
	"testing"
)

// TestFillMissingFieldsCacheDedup — when multiple specs share the same directory,
// a Diagnostic must be appended only once (seenDir deduplication).
func TestFillMissingFieldsCacheDedup(t *testing.T) {
	dir := t.TempDir()
	// A single file with a syntax error causes collectPackageTypes to emit one diag each call.
	bad := `package sample

!!!invalid!!!
`
	if err := os.WriteFile(filepath.Join(dir, "bad.go"), []byte(bad), 0644); err != nil {
		t.Fatal(err)
	}

	// Create two specs that reference the same dir.
	specs := []FuncSpec{
		{Name: "funcA", Package: "sample"},
		{Name: "funcB", Package: "sample"},
	}
	specDirs := []string{dir, dir}

	diags := fillMissingFields(specs, specDirs)
	if len(diags) != 1 {
		t.Fatalf("diags = %d, want 1 (dedup by dir): %v", len(diags), diags)
	}
}
