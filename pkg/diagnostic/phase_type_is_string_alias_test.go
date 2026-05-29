//ff:func feature=orchestrator type=test control=sequence
//ff:what Phase is string alias — string ↔ Phase 왕복 변환 검증
package diagnostic_test

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// TestPhase_TypeIsStringAlias verifies that Phase is convertible to and from string.
func TestPhase_TypeIsStringAlias(t *testing.T) {
	p := diagnostic.PhaseValidate
	s := string(p)
	if s != "validate" {
		t.Errorf("string(PhaseValidate): want %q, got %q", "validate", s)
	}

	// Reverse direction: string → Phase cast must also round-trip.
	custom := diagnostic.Phase("custom")
	if string(custom) != "custom" {
		t.Errorf("Phase(\"custom\"): round-trip failed, got %q", string(custom))
	}
}
