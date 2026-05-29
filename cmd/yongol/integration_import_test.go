//ff:func feature=cli type=test control=iteration dimension=1
//ff:what import file-source-happy — escrow fixture 기반 성공 시나리오

package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestIntegrationImport_FileSourceHappy runs `yongol import <local.yaml>
// <tmpdir>` on the bundled escrow OpenAPI fixture. Expects exit 0 and a
// generated .go file under the output dir. No network access — the source
// is a filesystem path, not a URL.
func TestIntegrationImport_FileSourceHappy(t *testing.T) {
	fixture := filepath.Join(repoRoot(t),
		"pkg", "external", "testdata", "escrow.openapi.yaml")
	if _, err := os.Stat(fixture); err != nil {
		t.Skipf("escrow fixture unavailable: %v", err)
	}
	outDir := t.TempDir()
	_, _, err := runCmd(t, "import", fixture, outDir)
	if err != nil {
		t.Fatalf("import happy: unexpected error: %v", err)
	}
	entries, readErr := os.ReadDir(outDir)
	if readErr != nil {
		t.Fatalf("read outDir: %v", readErr)
	}
	found := false
	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".go") {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected at least one .go file under %s, got %v", outDir, entries)
	}
}
