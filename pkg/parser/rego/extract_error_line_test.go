//ff:func feature=policy type=test control=sequence
//ff:what TestRegoHelpers — unit tests for the pure rego parser helper functions
package rego

import (
	"testing"
)

func TestExtractErrorLine(t *testing.T) {
	// Non-OPA error → 0.
	if got := extractErrorLine(errNotOPA{}); got != 0 {
		t.Errorf("non-OPA error → %d, want 0", got)
	}
	// nil error → 0 (type assertion fails).
	if got := extractErrorLine(nil); got != 0 {
		t.Errorf("nil error → %d, want 0", got)
	}
}
