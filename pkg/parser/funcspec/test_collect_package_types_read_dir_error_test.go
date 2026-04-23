//ff:func feature=funcspec type=test control=sequence
//ff:what collectPackageTypes — 파일 경로를 디렉토리로 받은 경우 1개 ReadDir 진단 발행

package funcspec

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestCollectPackageTypesReadDirError — passing a file path where a directory is expected
// must report the ReadDir error as exactly one Diagnostic.
func TestCollectPackageTypesReadDirError(t *testing.T) {
	base := t.TempDir()
	filePath := filepath.Join(base, "not_a_dir.txt")
	if err := os.WriteFile(filePath, []byte("hello"), 0644); err != nil {
		t.Fatal(err)
	}

	result, diags := collectPackageTypes(filePath)
	if len(diags) != 1 {
		t.Fatalf("diags = %d, want 1: %v", len(diags), diags)
	}
	d := diags[0]
	if d.File != filePath {
		t.Errorf("diag.File = %q, want %q", d.File, filePath)
	}
	if d.Phase != "parse" {
		t.Errorf("diag.Phase = %q, want parse", d.Phase)
	}
	if d.Level != "ERROR" {
		t.Errorf("diag.Level = %q, want ERROR", d.Level)
	}
	if !strings.Contains(d.Message, "cannot read funcspec type dir") {
		t.Errorf("diag.Message = %q, want ReadDir error message", d.Message)
	}
	if len(result) != 0 {
		t.Errorf("result should be empty on ReadDir error, got %v", keysOf(result))
	}
}
