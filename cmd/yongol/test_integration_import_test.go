//ff:func feature=cli type=test control=iteration dimension=1
//ff:what test: import 서브커맨드 end-to-end 3 케이스 (file-source-happy / missing-args / bad-source)

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

// TestIntegrationImport_MissingArgs verifies that `import <source>` (outDir
// missing) yields exit 2 via *usageError — consistent with the other
// subcommands that wrap cobra.ExactArgs through usageArgs.
func TestIntegrationImport_MissingArgs(t *testing.T) {
	_, _, err := runCmd(t, "import", "some-source.yaml")
	if err == nil {
		t.Fatal("expected usage error for missing outDir, got nil")
	}
	if !isUsageError(err) {
		t.Fatalf("expected *usageError (exit 2), got %T: %v", err, err)
	}
}

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
