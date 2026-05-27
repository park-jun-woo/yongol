//ff:func feature=cli type=test control=sequence
//ff:what agentCmd test — missing args, invalid model, missing dir

package main

import (
	"strings"
	"testing"
)

func TestAgentCmd(t *testing.T) {
	t.Run("MissingArgs", func(t *testing.T) {
		_, _, err := runCmd(t, "agent")
		if err == nil {
			t.Fatal("expected usage error for missing arg, got nil")
		}
		if !isUsageError(err) {
			t.Fatalf("expected *usageError (exit 2), got %T: %v", err, err)
		}
	})
	t.Run("InvalidModel", func(t *testing.T) {
		_, _, err := runCmd(t, "agent", "--model", "no-colon", "/tmp/fake-specs")
		if err == nil {
			t.Fatal("expected error for invalid model flag, got nil")
		}
		if !isUsageError(err) {
			t.Fatalf("expected *usageError (exit 2), got %T: %v", err, err)
		}
	})
	t.Run("MissingDir", func(t *testing.T) {
		_, stderr, err := runCmd(t, "agent", "/tmp/this-dir-should-not-exist-yongol-test")
		if err == nil {
			t.Fatal("expected error for missing dir, got nil")
		}
		_ = stderr
		// The error should be a runtime error (not usage), since the arg count is correct.
		if isUsageError(err) {
			t.Fatal("expected runtime error (exit 1), not usage error")
		}
	})
	t.Run("WarningBanner", func(t *testing.T) {
		_, stderr, _ := runCmd(t, "agent", "/tmp/this-dir-should-not-exist-yongol-test")
		if !strings.Contains(stderr, "실험 버전") {
			t.Errorf("expected warning banner in stderr, got: %s", stderr)
		}
	})
}
