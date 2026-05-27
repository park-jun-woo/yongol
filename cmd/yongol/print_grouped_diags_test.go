//ff:func feature=cli type=test control=sequence
//ff:what printGroupedDiags test — operationID 별 진단 그룹 출력 검증

package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestPrintGroupedDiags(t *testing.T) {
	t.Run("FiltersAndFormats", func(t *testing.T) {
		diags := []diagnostic.Diagnostic{
			{OperationID: "CreateWorkflow", Level: diagnostic.LevelError, Message: "S-01: missing arg", File: "specs/ssac/workflow.yaml", Line: 10},
			{OperationID: "DeleteWorkflow", Level: diagnostic.LevelError, Message: "S-02: unused var"},
			{OperationID: "CreateWorkflow", Level: diagnostic.LevelWarning, Message: "S-03: naming hint"},
		}
		var buf bytes.Buffer
		printGroupedDiags(&buf, "CreateWorkflow", diags, "specs/")
		out := buf.String()

		if !strings.Contains(out, "CreateWorkflow (2 errors)") {
			t.Errorf("expected 2 errors for CreateWorkflow, got: %q", out)
		}
		if !strings.Contains(out, "[ERROR] S-01") {
			t.Errorf("expected ERROR level, got: %q", out)
		}
		if !strings.Contains(out, "[WARNING] S-03") {
			t.Errorf("expected WARNING level, got: %q", out)
		}
		if strings.Contains(out, "S-02") {
			t.Errorf("should not include DeleteWorkflow diag, got: %q", out)
		}
		if !strings.Contains(out, "workflow.yaml:10") {
			t.Errorf("expected file:line, got: %q", out)
		}
	})

	t.Run("NoMatches", func(t *testing.T) {
		var buf bytes.Buffer
		printGroupedDiags(&buf, "MissingOp", nil, "specs/")
		out := buf.String()
		if !strings.Contains(out, "MissingOp (0 errors)") {
			t.Errorf("expected 0 errors, got: %q", out)
		}
	})
}
