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
		// A minimal but fully valid OpenAPI 3.0 document so external.Generate
		// succeeds and the RunE returns nil (success path).
		content := "" +
			"openapi: \"3.0.0\"\n" +
			"info:\n" +
			"  title: Test\n" +
			"  version: \"1.0\"\n" +
			"paths: {}\n"
		if err := os.WriteFile(spec, []byte(content), 0o644); err != nil {
			t.Fatal(err)
		}
		outDir := t.TempDir()
		_, _, err := runCmd(t, "import", spec, outDir)
		if err != nil {
			t.Fatalf("expected import success for valid spec, got: %v", err)
		}
		// The generated Go file should exist in the output dir.
		entries, derr := os.ReadDir(outDir)
		if derr != nil {
			t.Fatalf("read output dir: %v", derr)
		}
		if len(entries) == 0 {
			t.Fatal("expected at least one generated file in output dir")
		}
	})
}
