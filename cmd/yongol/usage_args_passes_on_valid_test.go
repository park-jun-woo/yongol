//ff:func feature=cli type=test control=sequence
//ff:what usageArgs passes on valid — 정상 인자는 래핑 없이 통과

package main

import (
	"testing"

	"github.com/spf13/cobra"
)

// TestUsageArgsPassesOnValid verifies that usageArgs returns nil when the
// underlying validator accepts the args (no wrapping on success).
func TestUsageArgsPassesOnValid(t *testing.T) {
	wrapped := usageArgs(cobra.ExactArgs(2))
	cmd := &cobra.Command{Use: "dummy"}
	if err := wrapped(cmd, []string{"a", "b"}); err != nil {
		t.Fatalf("expected nil for valid 2 args, got %v", err)
	}
}
