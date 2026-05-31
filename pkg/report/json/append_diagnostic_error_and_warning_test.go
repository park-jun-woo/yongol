//ff:func feature=report type=test control=sequence topic=json
//ff:what TestAppendDiagnostic — ERROR/WARNING 는 누적+카운트, 그 외 level 은 무시 검증
package json

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

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
