//ff:func feature=cli type=test control=sequence
//ff:what TestValidateCmd — validateCmd 서브커맨드 기본/포맷/에러 검증

package main

import (
	"bytes"
	"testing"
)

func TestValidateCmd(t *testing.T) {
	t.Run("ValidSpecs", func(t *testing.T) {
		cmd := validateCmd()
		var buf bytes.Buffer
		cmd.SetOut(&buf)
		cmd.SetErr(&buf)
		cmd.SetArgs([]string{"../../examples/zenflow/opus4_7/specs"})
		err := cmd.Execute()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("JSONFormat", func(t *testing.T) {
		cmd := validateCmd()
		var buf bytes.Buffer
		cmd.SetOut(&buf)
		cmd.SetErr(&buf)
		cmd.SetArgs([]string{"--format", "json", "../../examples/zenflow/opus4_7/specs"})
		err := cmd.Execute()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("SARIFFormat", func(t *testing.T) {
		cmd := validateCmd()
		var buf bytes.Buffer
		cmd.SetOut(&buf)
		cmd.SetErr(&buf)
		cmd.SetArgs([]string{"--format", "sarif", "../../examples/zenflow/opus4_7/specs"})
		err := cmd.Execute()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("InvalidFormat", func(t *testing.T) {
		cmd := validateCmd()
		var buf bytes.Buffer
		cmd.SetOut(&buf)
		cmd.SetErr(&buf)
		cmd.SetArgs([]string{"--format", "xml", "../../examples/zenflow/opus4_7/specs"})
		err := cmd.Execute()
		if err == nil {
			t.Fatal("expected error for invalid format")
		}
	})

	t.Run("InvalidSpecsDir", func(t *testing.T) {
		cmd := validateCmd()
		var buf bytes.Buffer
		cmd.SetOut(&buf)
		cmd.SetErr(&buf)
		cmd.SetArgs([]string{"/nonexistent/specs"})
		err := cmd.Execute()
		if err == nil {
			t.Fatal("expected error for nonexistent specs dir")
		}
	})

	t.Run("ParseErrorJSON", func(t *testing.T) {
		tmpDir := t.TempDir()
		// A malformed manifest.yaml is detected then fails to parse, driving the
		// ParseDiagnostics branch's JSON sub-path (wrapParseAsReport + printReport).
		if err := writeTestFile(tmpDir, "manifest.yaml", "::: not valid yaml :::\n\t- broken"); err != nil {
			t.Fatal(err)
		}
		cmd := validateCmd()
		var buf bytes.Buffer
		cmd.SetOut(&buf)
		cmd.SetErr(&buf)
		cmd.SetArgs([]string{"--format", "json", tmpDir})
		err := cmd.Execute()
		if err == nil {
			t.Fatal("expected parse error, got nil")
		}
		if err.Error() != "parse failed" {
			t.Fatalf("expected 'parse failed', got: %v", err)
		}
	})

	t.Run("ParseErrorMD", func(t *testing.T) {
		tmpDir := t.TempDir()
		// Default (md) format drives the else branch (printParseErrors).
		if err := writeTestFile(tmpDir, "manifest.yaml", "::: not valid yaml :::\n\t- broken"); err != nil {
			t.Fatal(err)
		}
		cmd := validateCmd()
		var buf bytes.Buffer
		cmd.SetOut(&buf)
		cmd.SetErr(&buf)
		cmd.SetArgs([]string{tmpDir})
		err := cmd.Execute()
		if err == nil {
			t.Fatal("expected parse error, got nil")
		}
		if err.Error() != "parse failed" {
			t.Fatalf("expected 'parse failed', got: %v", err)
		}
	})

	t.Run("WithArtsDir", func(t *testing.T) {
		cmd := validateCmd()
		var buf bytes.Buffer
		cmd.SetOut(&buf)
		cmd.SetErr(&buf)
		cmd.SetArgs([]string{"../../examples/zenflow/opus4_7/specs", "/tmp/nonexistent-arts"})
		err := cmd.Execute()
		// Should still succeed — arts dir is optional and contract check skips on missing.
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})
}
