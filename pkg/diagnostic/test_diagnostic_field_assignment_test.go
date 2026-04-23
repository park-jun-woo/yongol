//ff:func feature=orchestrator type=test control=sequence
//ff:what Diagnostic field assignment — 각 필드가 할당·비교 가능한지 검증
package diagnostic_test

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// TestDiagnostic_FieldAssignment verifies that each field can be assigned and compared.
func TestDiagnostic_FieldAssignment(t *testing.T) {
	d := diagnostic.Diagnostic{
		File:    "spec.yaml",
		Line:    42,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: "X-1: bad thing",
		Advice:  "fix it",
	}

	if d.File != "spec.yaml" {
		t.Errorf("File: want %q, got %q", "spec.yaml", d.File)
	}
	if d.Line != 42 {
		t.Errorf("Line: want 42, got %d", d.Line)
	}
	if d.Phase != diagnostic.PhaseValidate {
		t.Errorf("Phase: want %q, got %q", diagnostic.PhaseValidate, d.Phase)
	}
	if d.Level != diagnostic.LevelError {
		t.Errorf("Level: want %q, got %q", diagnostic.LevelError, d.Level)
	}
	if d.Message != "X-1: bad thing" {
		t.Errorf("Message: want %q, got %q", "X-1: bad thing", d.Message)
	}
	if d.Advice != "fix it" {
		t.Errorf("Advice: want %q, got %q", "fix it", d.Advice)
	}
}
