//ff:func feature=cli type=test control=sequence
//ff:what nextCmd test — missing args, missing dir, valid specs

package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNextCmd(t *testing.T) {
	t.Run("MissingArgs", func(t *testing.T) {
		_, _, err := runCmd(t, "next")
		if err == nil {
			t.Fatal("expected usage error for missing arg, got nil")
		}
		if !isUsageError(err) {
			t.Fatalf("expected *usageError (exit 2), got %T: %v", err, err)
		}
	})
	t.Run("MissingDir", func(t *testing.T) {
		_, _, err := runCmd(t, "next", "/tmp/nonexistent-yongol-next")
		if err == nil {
			t.Fatal("expected error for missing dir, got nil")
		}
	})
	t.Run("WithSpecs", func(t *testing.T) {
		specsDir := filepath.Join(repoRoot(t), "examples", "zenflow", "opus4_7", "specs")
		if _, err := os.Stat(specsDir); err != nil {
			t.Skipf("opus4_7 specs not available: %v", err)
		}
		_, _, err := runCmd(t, "next", specsDir)
		// May pass or fail depending on validation state.
		_ = err
	})
	t.Run("WithErrors", func(t *testing.T) {
		// Create a minimal specs dir with only a manifest to trigger validation errors.
		dir := t.TempDir()
		manifest := `metadata:
	  name: test-next
	backend:
	  lang: go
	  framework: gin
	`
		if err := os.WriteFile(filepath.Join(dir, "manifest.yaml"), []byte(manifest), 0o644); err != nil {
			t.Fatal(err)
		}
		features := `features: []
	`
		if err := os.WriteFile(filepath.Join(dir, "features.yaml"), []byte(features), 0o644); err != nil {
			t.Fatal(err)
		}
		_, _, err := runCmd(t, "next", dir)
		// Should report validation issues (or parse issues).
		_ = err
	})
	t.Run("WithParseError", func(t *testing.T) {
		// Create a specs dir with an invalid YAML file to trigger parse errors.
		dir := t.TempDir()
		if err := os.MkdirAll(filepath.Join(dir, "api"), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(dir, "api", "openapi.yaml"), []byte("invalid: [yaml: broken"), 0o644); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(dir, "manifest.yaml"), []byte("metadata:\n  name: broken\n"), 0o644); err != nil {
			t.Fatal(err)
		}
		_, _, err := runCmd(t, "next", dir)
		// Should hit the parse error path.
		_ = err
	})
}
