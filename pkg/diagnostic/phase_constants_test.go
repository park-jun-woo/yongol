//ff:func feature=orchestrator type=test control=iteration dimension=1
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
