//ff:func feature=cli type=test control=sequence
//ff:what printParseErrors test — Parse Errors 섹션 출력 검증

package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestPrintParseErrors(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		var buf bytes.Buffer
		printParseErrors(&buf, nil)
		out := buf.String()
		if !strings.Contains(out, "Parse Errors") {
			t.Errorf("expected header, got: %q", out)
		}
		if !strings.Contains(out, "0 parse errors") {
			t.Errorf("expected 0 count, got: %q", out)
		}
	})

	t.Run("GroupedByFile", func(t *testing.T) {
		diags := []diagnostic.Diagnostic{
			{File: "specs/ddl/b.sql", Level: diagnostic.LevelError, Message: "D-1: bad", Line: 5},
			{File: "specs/ddl/a.sql", Level: diagnostic.LevelError, Message: "D-2: missing"},
			{File: "specs/ddl/b.sql", Level: diagnostic.LevelWarning, Message: "D-3: warn", Line: 10},
		}
		var buf bytes.Buffer
		printParseErrors(&buf, diags)
		out := buf.String()

		if !strings.Contains(out, "3 parse errors") {
			t.Errorf("expected 3 errors, got: %q", out)
		}
		// a.sql should come before b.sql.
		aIdx := strings.Index(out, "a.sql")
		bIdx := strings.Index(out, "b.sql")
		if aIdx < 0 || bIdx < 0 || aIdx > bIdx {
			t.Errorf("expected a.sql before b.sql, got: %q", out)
		}
		if !strings.Contains(out, "(line 5)") {
			t.Errorf("expected line 5, got: %q", out)
		}
	})
}
