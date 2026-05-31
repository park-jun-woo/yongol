//ff:func feature=cli type=test control=sequence
//ff:what TestStatusCmd — statusCmd 서브커맨드 기본 동작 검증
package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestStatusCmd(t *testing.T) {
	t.Run("ValidSpecs", func(t *testing.T) {
		cmd := statusCmd()
		var buf bytes.Buffer
		cmd.SetOut(&buf)
		cmd.SetErr(&buf)
		cmd.SetArgs([]string{"../../examples/zenflow/opus4_7/specs"})
		err := cmd.Execute()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		out := buf.String()
		if !strings.Contains(out, "SSOT Summary") {
			t.Error("expected 'SSOT Summary' in output")
		}
	})

	t.Run("InvalidSpecsDir", func(t *testing.T) {
		cmd := statusCmd()
		var buf bytes.Buffer
		cmd.SetOut(&buf)
		cmd.SetErr(&buf)
		cmd.SetArgs([]string{"/nonexistent/path/to/specs"})
		err := cmd.Execute()
		if err == nil {
			t.Fatal("expected error for invalid specs dir")
		}
	})

	t.Run("NoArgs", func(t *testing.T) {
		cmd := statusCmd()
		var buf bytes.Buffer
		cmd.SetOut(&buf)
		cmd.SetErr(&buf)
		cmd.SetArgs([]string{})
		err := cmd.Execute()
		if err == nil {
			t.Fatal("expected error for no args")
		}
	})

	t.Run("WithExistingArtsDir", func(t *testing.T) {
		artsDir := t.TempDir()
		cmd := statusCmd()
		var buf bytes.Buffer
		cmd.SetOut(&buf)
		cmd.SetErr(&buf)
		cmd.SetArgs([]string{"../../examples/zenflow/opus4_7/specs", artsDir})
		err := cmd.Execute()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	t.Run("ParseError", func(t *testing.T) {
		// A malformed manifest.yaml is detected as an SSOT and fails to parse,
		// driving the ParseDiagnostics branch (printParseErrors + "parse failed").
		tmpDir := t.TempDir()
		if err := writeTestFile(tmpDir, "manifest.yaml", "::: not valid yaml :::\n\t- broken"); err != nil {
			t.Fatal(err)
		}
		cmd := statusCmd()
		var buf bytes.Buffer
		cmd.SetOut(&buf)
		cmd.SetErr(&buf)
		cmd.SetArgs([]string{tmpDir})
		err := cmd.Execute()
		if err == nil {
			t.Fatal("expected parse error, got nil")
		}
		if !strings.Contains(err.Error(), "parse failed") {
			t.Fatalf("expected 'parse failed' error, got: %v", err)
		}
	})

	t.Run("WithArtsDir", func(t *testing.T) {
		cmd := statusCmd()
		var buf bytes.Buffer
		cmd.SetOut(&buf)
		cmd.SetErr(&buf)
		cmd.SetArgs([]string{"../../examples/zenflow/opus4_7/specs", "/nonexistent/arts"})
		err := cmd.Execute()
		if err != nil {
			t.Fatalf("unexpected error: %v (status should not fail on missing arts)", err)
		}
		out := buf.String()
		if !strings.Contains(out, "no artifacts") {
			t.Error("expected 'no artifacts' in output for missing arts dir")
		}
	})
}
