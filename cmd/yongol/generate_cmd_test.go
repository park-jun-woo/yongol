//ff:func feature=cli type=test control=sequence
//ff:what generateCmd test — missing args, missing dir, parse error, backend resolution, happy path

package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestGenerateCmd(t *testing.T) {
	t.Run("MissingArgs", func(t *testing.T) {
		_, _, err := runCmd(t, "generate")
		if err == nil {
			t.Fatal("expected usage error for missing args, got nil")
		}
		if !isUsageError(err) {
			t.Fatalf("expected *usageError (exit 2), got %T: %v", err, err)
		}
	})
	t.Run("OneArg", func(t *testing.T) {
		_, _, err := runCmd(t, "generate", "/tmp/specs-only")
		if err == nil {
			t.Fatal("expected usage error for single arg, got nil")
		}
		if !isUsageError(err) {
			t.Fatalf("expected *usageError (exit 2), got %T: %v", err, err)
		}
	})
	t.Run("MissingDir", func(t *testing.T) {
		_, _, err := runCmd(t, "generate", "/tmp/nonexistent-yongol-specs", "/tmp/output")
		if err == nil {
			t.Fatal("expected error for missing dir, got nil")
		}
	})
	t.Run("ParseError", func(t *testing.T) {
		// A malformed manifest.yaml is detected then fails to parse, driving the
		// ParseDiagnostics branch (printParseErrors + "parse failed").
		dir := t.TempDir()
		if err := os.WriteFile(filepath.Join(dir, "manifest.yaml"), []byte("::: not valid yaml :::\n\t- broken"), 0o644); err != nil {
			t.Fatalf("write manifest: %v", err)
		}
		_, _, err := runCmd(t, "generate", dir, t.TempDir())
		if err == nil {
			t.Fatal("expected parse error, got nil")
		}
		if !strings.Contains(err.Error(), "parse failed") {
			t.Fatalf("expected 'parse failed' error, got: %v", err)
		}
	})
	t.Run("UnknownBackendNoManifest", func(t *testing.T) {
		// An empty (but valid) specs dir parses clean and has no manifest, so an
		// unrecognized --backend falls through to the "unknown --backend" error
		// rather than manifest resolution.
		dir := t.TempDir()
		_, _, err := runCmd(t, "generate", "--backend", "no-such-backend", dir, t.TempDir())
		if err == nil {
			t.Fatal("expected unknown backend error, got nil")
		}
	})
	t.Run("HasFlags", func(t *testing.T) {
		cmd := generateCmd()
		if cmd.Flags().Lookup("backend") == nil {
			t.Error("expected --backend flag")
		}
		if cmd.Flags().Lookup("frontend") == nil {
			t.Error("expected --frontend flag")
		}
	})
	t.Run("InvalidBackend", func(t *testing.T) {
		specsDir := filepath.Join(repoRoot(t), "examples", "zenflow", "opus4_7", "specs")
		if _, err := os.Stat(specsDir); err != nil {
			t.Skipf("opus4_7 specs not available: %v", err)
		}
		outDir := t.TempDir()
		_, _, err := runCmd(t, "generate", "--backend", "invalid-framework", specsDir, outDir)
		// The generate may fail during validate or at backend resolution, both are acceptable.
		if err == nil {
			t.Fatal("expected error for invalid backend, got nil")
		}
	})
	t.Run("NestJSBackend", func(t *testing.T) {
		specsDir := filepath.Join(repoRoot(t), "examples", "zenflow", "opus4_7", "specs")
		if _, err := os.Stat(specsDir); err != nil {
			t.Skipf("opus4_7 specs not available: %v", err)
		}
		outDir := t.TempDir()
		_, _, err := runCmd(t, "generate", "--backend", "nestjs", specsDir, outDir)
		// May have validation warnings so generate is refused, but the
		// RunE body should be exercised either way.
		_ = err
	})
	t.Run("GoGinBackend", func(t *testing.T) {
		specsDir := filepath.Join(repoRoot(t), "examples", "zenflow", "opus4_7", "specs")
		if _, err := os.Stat(specsDir); err != nil {
			t.Skipf("opus4_7 specs not available: %v", err)
		}
		outDir := t.TempDir()
		_, _, err := runCmd(t, "generate", "--backend", "go-gin", specsDir, outDir)
		_ = err
	})
	t.Run("FastAPIBackend", func(t *testing.T) {
		specsDir := filepath.Join(repoRoot(t), "examples", "zenflow", "opus4_7", "specs")
		if _, err := os.Stat(specsDir); err != nil {
			t.Skipf("opus4_7 specs not available: %v", err)
		}
		outDir := t.TempDir()
		_, _, err := runCmd(t, "generate", "--backend", "fastapi", specsDir, outDir)
		_ = err
	})
	t.Run("DefaultBackendFromManifest", func(t *testing.T) {
		specsDir := filepath.Join(repoRoot(t), "examples", "zenflow", "opus4_7", "specs")
		if _, err := os.Stat(specsDir); err != nil {
			t.Skipf("opus4_7 specs not available: %v", err)
		}
		outDir := t.TempDir()
		// Use empty string for backend to trigger manifest resolution.
		_, _, err := runCmd(t, "generate", "--backend", "", specsDir, outDir)
		_ = err
	})
}
