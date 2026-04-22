//ff:func feature=cli type=test control=sequence
//ff:what test: TestUsageArgsAndErrorsAs — usageArgs wraps/passes + errors.As 매칭 검증

package main

import (
	"errors"
	"testing"

	"github.com/spf13/cobra"
)

// TestUsageArgsWrapsFailure verifies that when the underlying cobra
// PositionalArgs validator fails, usageArgs wraps the error in *usageError.
func TestUsageArgsWrapsFailure(t *testing.T) {
	wrapped := usageArgs(cobra.ExactArgs(2))
	cmd := &cobra.Command{Use: "dummy"}
	err := wrapped(cmd, []string{"only-one"})
	if err == nil {
		t.Fatal("expected error for 1 arg when 2 required, got nil")
	}
	var ue *usageError
	if !errors.As(err, &ue) {
		t.Fatalf("expected *usageError, got %T: %v", err, err)
	}
}

// TestUsageArgsPassesOnValid verifies that usageArgs returns nil when the
// underlying validator accepts the args (no wrapping on success).
func TestUsageArgsPassesOnValid(t *testing.T) {
	wrapped := usageArgs(cobra.ExactArgs(2))
	cmd := &cobra.Command{Use: "dummy"}
	if err := wrapped(cmd, []string{"a", "b"}); err != nil {
		t.Fatalf("expected nil for valid 2 args, got %v", err)
	}
}

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
