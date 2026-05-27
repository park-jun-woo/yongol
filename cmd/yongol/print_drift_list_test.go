//ff:func feature=cli type=test control=sequence
//ff:what printDriftList test — Contract Drift 섹션 출력 검증

package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestPrintDriftList(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		var buf bytes.Buffer
		printDriftList(&buf, nil)
		out := buf.String()
		if !strings.Contains(out, "Contract Drift (0)") {
			t.Errorf("expected header with 0 count, got: %q", out)
		}
	})

	t.Run("GroupedByFile", func(t *testing.T) {
		diags := []diagnostic.Diagnostic{
			{File: "arts/backend/svc.go", Message: "PRV-01: missing hash"},
			{File: "arts/backend/handler.go", Message: "PRV-02: body changed"},
			{File: "arts/backend/svc.go", Message: "PRV-01: stale contract"},
		}
		var buf bytes.Buffer
		printDriftList(&buf, diags)
		out := buf.String()

		if !strings.Contains(out, "Contract Drift (3)") {
			t.Errorf("expected count 3, got: %q", out)
		}
		// handler.go comes before svc.go alphabetically.
		hIdx := strings.Index(out, "handler.go")
		sIdx := strings.Index(out, "svc.go")
		if hIdx < 0 || sIdx < 0 || hIdx > sIdx {
			t.Errorf("expected handler.go before svc.go, got: %q", out)
		}
		if !strings.Contains(out, "PRV-01: missing hash") {
			t.Errorf("expected message present, got: %q", out)
		}
	})
}
