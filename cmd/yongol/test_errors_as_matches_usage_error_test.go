//ff:func feature=cli type=test control=sequence
//ff:what errors.As — manually 생성한 *usageError 가 main.go unwrap 로직과 매치

package main

import (
	"errors"
	"testing"
)

// TestErrorsAsMatchesUsageError verifies errors.As unwraps a manually
// constructed *usageError the same way main.go's exit-code branch does.
func TestErrorsAsMatchesUsageError(t *testing.T) {
	inner := errors.New("unknown flag: --nope")
	wrapped := &usageError{err: inner}

	var ue *usageError
	if !errors.As(wrapped, &ue) {
		t.Fatalf("errors.As did not match *usageError; got false")
	}
	if ue.Error() != inner.Error() {
		t.Errorf("Error() mismatch: got %q, want %q", ue.Error(), inner.Error())
	}
	if !errors.Is(wrapped, inner) {
		t.Errorf("errors.Is did not unwrap to inner; Unwrap() chain broken")
	}
}
