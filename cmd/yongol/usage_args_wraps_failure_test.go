//ff:func feature=cli type=test control=sequence
//ff:what usageArgs wraps failure — cobra PositionalArgs 실패 시 *usageError 래핑

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
