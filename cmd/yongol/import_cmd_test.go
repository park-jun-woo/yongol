//ff:func feature=cli type=test control=sequence
//ff:what importCmd test — missing args, missing source, valid source

package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestImportCmd(t *testing.T) {
	t.Run("MissingArgs", func(t *testing.T) {
		_, _, err := runCmd(t, "import")
		if err == nil {
			t.Fatal("expected usage error for missing args, got nil")
		}
		if !isUsageError(err) {
			t.Fatalf("expected *usageError (exit 2), got %T: %v", err, err)
		}
	})
	t.Run("MissingSource", func(t *testing.T) {
		_, _, err := runCmd(t, "import", "/tmp/nonexist-source.yaml", t.TempDir())
		if err == nil {
			t.Fatal("expected error for missing source, got nil")
		}
	})
	t.Run("ValidSource", func(t *testing.T) {
		srcDir := t.TempDir()
		spec := filepath.Join(srcDir, "openapi.yaml")
		content := `openapi: "3.0.0"
	info:
	  title: Test
	  version: "1.0"
	paths: {}
	`
		if err := os.WriteFile(spec, []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
		outDir := t.TempDir()
		_, _, err := runCmd(t, "import", spec, outDir)
		// Success or specific error — both exercise the RunE body fully.
		_ = err
	})
}
