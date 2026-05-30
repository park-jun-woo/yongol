//ff:func feature=orchestrator type=test control=sequence
//ff:what TestLogBuiltinLoadIssues — 빈/비빈 diags 모두에서 slog.Warn 호출 검증

package yongol

import (
	"bytes"
	"log/slog"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestLogBuiltinLoadIssues(t *testing.T) {
	var buf bytes.Buffer
	orig := slog.Default()
	slog.SetDefault(slog.New(slog.NewTextHandler(&buf, nil)))
	t.Cleanup(func() { slog.SetDefault(orig) })

	// Empty diags → first message stays empty.
	logBuiltinLoadIssues("/pkg/root", nil)
	if !strings.Contains(buf.String(), "diag_count=0") {
		t.Fatalf("expected diag_count=0 in log, got %q", buf.String())
	}

	buf.Reset()
	// Non-empty diags → first message included.
	logBuiltinLoadIssues("/pkg/root", []diagnostic.Diagnostic{
		{Message: "boom"},
		{Message: "second"},
	})
	out := buf.String()
	if !strings.Contains(out, "diag_count=2") || !strings.Contains(out, "boom") {
		t.Fatalf("expected diag_count=2 and first message, got %q", out)
	}
}
