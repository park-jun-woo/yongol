//ff:func feature=orchestrator type=test control=sequence
//ff:what Phase constant value lock-in test
package diagnostic_test

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// TestPhaseConstants locks in the string values of Phase constants.
// It will fail intentionally on any value change — update the test alongside the constant.
func TestPhaseConstants(t *testing.T) {
	cases := []struct {
		name string
		got  diagnostic.Phase
		want string
	}{
		{"PhaseParse", diagnostic.PhaseParse, "parse"},
		{"PhaseValidate", diagnostic.PhaseValidate, "validate"},
	}

	for _, c := range cases {
		if string(c.got) != c.want {
			t.Errorf("%s value changed: want %q, got %q", c.name, c.want, string(c.got))
		}
	}
}

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
