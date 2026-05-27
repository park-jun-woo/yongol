//ff:func feature=cli type=test control=sequence
//ff:what printDiagFile test — 파일 경로+줄번호 출력 검증

package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestPrintDiagFile(t *testing.T) {
	t.Run("WithLine", func(t *testing.T) {
		var buf bytes.Buffer
		d := diagnostic.Diagnostic{File: "specs/ddl/foo.sql", Line: 42}
		printDiagFile(&buf, d)
		out := buf.String()
		if !strings.Contains(out, "foo.sql:42") {
			t.Errorf("expected file:line, got: %q", out)
		}
	})

	t.Run("WithoutLine", func(t *testing.T) {
		var buf bytes.Buffer
		d := diagnostic.Diagnostic{File: "specs/ddl/bar.sql"}
		printDiagFile(&buf, d)
		out := buf.String()
		if !strings.Contains(out, "bar.sql") {
			t.Errorf("expected file only, got: %q", out)
		}
		// When Line==0 the output should not have ":NNN" after the filename.
		if strings.Contains(out, "bar.sql:") {
			t.Errorf("expected no line number suffix when Line=0, got: %q", out)
		}
	})

	t.Run("EmptyFile", func(t *testing.T) {
		var buf bytes.Buffer
		d := diagnostic.Diagnostic{}
		printDiagFile(&buf, d)
		if buf.Len() != 0 {
			t.Errorf("expected no output for empty file, got: %q", buf.String())
		}
	})
}
