//ff:type feature=orchestrator type=test
//ff:what Regression test that locks in the zero value of the Diagnostic struct
package diagnostic_test

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// TestDiagnostic_ZeroValue locks in the expected zero value of Diagnostic.
func TestDiagnostic_ZeroValue(t *testing.T) {
	var d diagnostic.Diagnostic

	if d.File != "" {
		t.Errorf("File zero: want %q, got %q", "", d.File)
	}
	if d.Line != 0 {
		t.Errorf("Line zero: want 0, got %d", d.Line)
	}
	if d.Phase != diagnostic.Phase("") {
		t.Errorf("Phase zero: want empty, got %q", d.Phase)
	}
	if d.Level != diagnostic.Level("") {
		t.Errorf("Level zero: want empty, got %q", d.Level)
	}
	if d.Message != "" {
		t.Errorf("Message zero: want %q, got %q", "", d.Message)
	}
	if d.Advice != "" {
		t.Errorf("Advice zero: want %q, got %q", "", d.Advice)
	}
}

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

// TestDiagnostic_DeepEqual verifies that two Diagnostics with the same values compare equal via reflect.DeepEqual.
func TestDiagnostic_DeepEqual(t *testing.T) {
	a := diagnostic.Diagnostic{
		File:    "a.yaml",
		Line:    1,
		Phase:   diagnostic.PhaseParse,
		Level:   diagnostic.LevelWarning,
		Message: "m",
		Advice:  "adv",
	}
	b := diagnostic.Diagnostic{
		File:    "a.yaml",
		Line:    1,
		Phase:   diagnostic.PhaseParse,
		Level:   diagnostic.LevelWarning,
		Message: "m",
		Advice:  "adv",
	}

	if !reflect.DeepEqual(a, b) {
		t.Errorf("DeepEqual: want equal, got diff\n  a=%+v\n  b=%+v", a, b)
	}

	b.Line = 2
	if reflect.DeepEqual(a, b) {
		t.Errorf("DeepEqual: want diff after Line change, still equal")
	}
}
