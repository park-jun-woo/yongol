//ff:func feature=cli type=test control=sequence
//ff:what chainCmd test — missing args, missing dir, parse error, happy path, unknown op

package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestChainCmd(t *testing.T) {
	t.Run("MissingArgs", func(t *testing.T) {
		_, _, err := runCmd(t, "chain")
		if err == nil {
			t.Fatal("expected usage error for missing args, got nil")
		}
		if !isUsageError(err) {
			t.Fatalf("expected *usageError (exit 2), got %T: %v", err, err)
		}
	})
	t.Run("MissingDir", func(t *testing.T) {
		_, _, err := runCmd(t, "chain", "SomeOpID", "/tmp/this-dir-should-not-exist-yongol-chain")
		if err == nil {
			t.Fatal("expected error for missing dir, got nil")
		}
	})
	t.Run("ParseError", func(t *testing.T) {
		dir := t.TempDir()
		// A malformed manifest.yaml is detected as an SSOT but fails to parse,
		// driving the ParseDiagnostics branch (printParseErrors + "parse failed").
		if err := os.WriteFile(filepath.Join(dir, "manifest.yaml"), []byte("::: not valid yaml :::\n\t- broken"), 0o644); err != nil {
			t.Fatalf("write manifest: %v", err)
		}
		_, _, err := runCmd(t, "chain", "SomeOpID", dir)
		if err == nil {
			t.Fatal("expected parse error, got nil")
		}
		if !strings.Contains(err.Error(), "parse failed") {
			t.Fatalf("expected 'parse failed' error, got: %v", err)
		}
	})
	t.Run("Happy", func(t *testing.T) {
		specs := opus47SpecsDir(t)
		stdout, _, err := runCmd(t, "chain", "ExecuteWorkflow", specs)
		if err != nil {
			t.Fatalf("chain happy: unexpected error: %v", err)
		}
		if !strings.Contains(stdout, "Feature Chain:") {
			t.Errorf("expected Feature Chain header, got:\n%s", stdout)
		}
	})
	t.Run("UnknownOp", func(t *testing.T) {
		specs := opus47SpecsDir(t)
		_, _, err := runCmd(t, "chain", "NonExistentOperation99999", specs)
		if err == nil {
			t.Fatal("expected error for unknown operation ID, got nil")
		}
	})
}
