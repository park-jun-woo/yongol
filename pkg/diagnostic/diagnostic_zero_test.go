//ff:func feature=orchestrator type=test control=sequence
//ff:what Regression test that locks in the zero value of the Diagnostic struct
package diagnostic_test

import (
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
