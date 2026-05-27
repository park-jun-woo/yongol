//ff:func feature=cli type=test control=sequence
//ff:what printNextDiag test — next 형식 단일 진단 출력 검증

package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/spf13/cobra"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestPrintNextDiag(t *testing.T) {
	t.Run("WithFileLine", func(t *testing.T) {
		var buf bytes.Buffer
		cmd := &cobra.Command{}
		cmd.SetOut(&buf)

		d := diagnostic.Diagnostic{
			Level:   diagnostic.LevelError,
			Message: "S-01: missing arg",
			File:    "specs/ssac/workflow.yaml",
			Line:    10,
		}
		printNextDiag(cmd, d, "specs/")
		out := buf.String()

		if !strings.Contains(out, "[ERROR] S-01: missing arg") {
			t.Errorf("expected error message, got: %q", out)
		}
		if !strings.Contains(out, "workflow.yaml:10") {
			t.Errorf("expected file:line, got: %q", out)
		}
		if !strings.Contains(out, "yongol next specs/") {
			t.Errorf("expected next instruction, got: %q", out)
		}
	})

	t.Run("WithoutLine", func(t *testing.T) {
		var buf bytes.Buffer
		cmd := &cobra.Command{}
		cmd.SetOut(&buf)

		d := diagnostic.Diagnostic{
			Level:   diagnostic.LevelError,
			Message: "O-01: bad spec",
			File:    "specs/openapi.yaml",
		}
		printNextDiag(cmd, d, "specs/")
		out := buf.String()

		if !strings.Contains(out, "openapi.yaml") {
			t.Errorf("expected file, got: %q", out)
		}
		if strings.Contains(out, "openapi.yaml:") {
			t.Errorf("expected no line number suffix, got: %q", out)
		}
	})

	t.Run("NoFile", func(t *testing.T) {
		var buf bytes.Buffer
		cmd := &cobra.Command{}
		cmd.SetOut(&buf)

		d := diagnostic.Diagnostic{
			Level:   diagnostic.LevelWarning,
			Message: "C-01: check",
		}
		printNextDiag(cmd, d, "specs/")
		out := buf.String()

		if strings.Contains(out, "file:") {
			t.Errorf("expected no file line, got: %q", out)
		}
	})
}
