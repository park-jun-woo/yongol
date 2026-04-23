//ff:func feature=cli type=test control=sequence
//ff:what import bad-source — 존재하지 않는 파일 시 read source 에러 (exit 1)

package main

import (
	"path/filepath"
	"strings"
	"testing"
)

// TestIntegrationImport_BadSource points at a non-existent file and expects
// exit 1 with an error message that surfaces the read-source failure. The
// path deliberately lives under t.TempDir() so no stray FS access occurs.
func TestIntegrationImport_BadSource(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "no-such-file.yaml")
	outDir := t.TempDir()
	_, _, err := runCmd(t, "import", missing, outDir)
	if err == nil {
		t.Fatal("expected error for missing source file, got nil")
	}
	if isUsageError(err) {
		t.Fatalf("bad-source must be exit 1, not usage error: %v", err)
	}
	if !strings.Contains(err.Error(), "import failed") {
		t.Errorf("expected error to wrap with `import failed`, got: %v", err)
	}
	if !strings.Contains(err.Error(), "read source") {
		t.Errorf("expected error to mention `read source`, got: %v", err)
	}
}
