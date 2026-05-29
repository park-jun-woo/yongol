//ff:func feature=report type=test control=selection topic=json
//ff:what TestAppendDiagnostic — ERROR/WARNING 는 누적+카운트, 그 외 level 은 무시 검증
package json

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// TestAppendDiagnostic_ErrorAndWarning covers the two appending branches:
// ERROR increments errors and WARNING increments warnings, both producing a
// Diagnostic entry with rule_id extracted and file rebased.
func TestAppendDiagnostic_ErrorAndWarning(t *testing.T) {
	doc := &Document{Diagnostics: []Diagnostic{}}

	appendDiagnostic(doc, diagnostic.Diagnostic{
		File:    "auth/login.ssac",
		Line:    15,
		Level:   diagnostic.LevelError,
		Message: "[S-27] foo not declared",
	}, "", "")

	appendDiagnostic(doc, diagnostic.Diagnostic{
		File:    "auth/update.ssac",
		Line:    8,
		Level:   diagnostic.LevelWarning,
		Message: "[S-36] stale response",
	}, "", "")

	if doc.Summary.Errors != 1 {
		t.Errorf("Errors: got %d, want 1", doc.Summary.Errors)
	}
	if doc.Summary.Warnings != 1 {
		t.Errorf("Warnings: got %d, want 1", doc.Summary.Warnings)
	}
	if len(doc.Diagnostics) != 2 {
		t.Fatalf("Diagnostics: got %d, want 2", len(doc.Diagnostics))
	}

	d0 := doc.Diagnostics[0]
	if d0.RuleID != "S-27" || d0.Level != "ERROR" || d0.Line != 15 {
		t.Errorf("diagnostics[0]: %+v", d0)
	}
	if d0.File != "auth/login.ssac" {
		t.Errorf("diagnostics[0].File: got %q, want auth/login.ssac", d0.File)
	}
	if d0.Message != "foo not declared" {
		t.Errorf("diagnostics[0].Message: got %q, want %q", d0.Message, "foo not declared")
	}
	if d0.Col != 0 {
		t.Errorf("diagnostics[0].Col: got %d, want 0", d0.Col)
	}

	d1 := doc.Diagnostics[1]
	if d1.RuleID != "S-36" || d1.Level != "WARNING" {
		t.Errorf("diagnostics[1]: %+v", d1)
	}
}

// TestAppendDiagnostic_LowSeverityIgnored covers the default branch: a
// non-error/non-warning diagnostic is dropped entirely (no entry, no count).
func TestAppendDiagnostic_LowSeverityIgnored(t *testing.T) {
	doc := &Document{Diagnostics: []Diagnostic{}}

	appendDiagnostic(doc, diagnostic.Diagnostic{
		File:    "x.ssac",
		Level:   diagnostic.Level("INFO"),
		Message: "[S-99] purely informational",
	}, "", "")

	if doc.Summary.Errors != 0 || doc.Summary.Warnings != 0 {
		t.Errorf("counts should be zero, got %+v", doc.Summary)
	}
	if len(doc.Diagnostics) != 0 {
		t.Errorf("diagnostics should be empty, got %d", len(doc.Diagnostics))
	}
}
