//ff:func feature=cli type=test control=sequence
//ff:what featuresCmd test — parent command returns subcommands

package main

import (
	"testing"
)

func TestFeaturesCmd(t *testing.T) {
	t.Run("NoSubcommand", func(t *testing.T) {
		_, _, err := runCmd(t, "features")
		// Running "features" without a subcommand should not error (cobra shows help)
		if err != nil {
			t.Fatalf("expected nil for features parent cmd, got: %v", err)
		}
	})
	t.Run("HasSubcommands", func(t *testing.T) {
		cmd := featuresCmd()
		if len(cmd.Commands()) < 2 {
			t.Errorf("expected at least 2 subcommands (add, remove), got %d", len(cmd.Commands()))
		}
	})
}
