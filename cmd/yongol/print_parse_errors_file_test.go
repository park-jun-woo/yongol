//ff:func feature=cli type=test control=sequence
//ff:what printParseErrorsFile test — 단일 파일의 진단 출력 검증

package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestPrintParseErrorsFile(t *testing.T) {
	t.Run("WithLine", func(t *testing.T) {
		diags := []diagnostic.Diagnostic{
			{Level: diagnostic.LevelError, Message: "D-1: bad col", Line: 5},
		}
		var buf bytes.Buffer
		printParseErrorsFile(&buf, "specs/ddl/users.sql", diags)
		out := buf.String()

		if !strings.Contains(out, "[specs/ddl/users.sql]") {
			t.Errorf("expected file header, got: %q", out)
		}
		if !strings.Contains(out, "(line 5)") {
			t.Errorf("expected line number, got: %q", out)
		}
		if !strings.Contains(out, "[ERROR]") {
			t.Errorf("expected ERROR level, got: %q", out)
		}
	})

	t.Run("WithoutLine", func(t *testing.T) {
		diags := []diagnostic.Diagnostic{
			{Level: diagnostic.LevelWarning, Message: "D-2: warn"},
		}
		var buf bytes.Buffer
		printParseErrorsFile(&buf, "specs/openapi.yaml", diags)
		out := buf.String()

		if strings.Contains(out, "(line") {
			t.Errorf("expected no line info, got: %q", out)
		}
		if !strings.Contains(out, "[WARNING]") {
			t.Errorf("expected WARNING level, got: %q", out)
		}
	})
}
