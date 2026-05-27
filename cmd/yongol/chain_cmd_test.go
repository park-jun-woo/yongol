//ff:func feature=cli type=test control=sequence
//ff:what chainCmd test — missing args, missing dir, happy path, unknown op

package main

import (
	"strings"
	"testing"
)

func TestChainCmd(t *testing.T) {
	t.Run("MissingArgs", func(t *testing.T) {
		_, _, err := runCmd(t, "chain")
		if err == nil {
			t.Fatal("expected usage error for missing args, got nil")
		}
		if !isUsageError(err) {
			t.Fatalf("expected *usageError (exit 2), got %T: %v", err, err)
		}
	})
	t.Run("MissingDir", func(t *testing.T) {
		_, _, err := runCmd(t, "chain", "SomeOpID", "/tmp/this-dir-should-not-exist-yongol-chain")
		if err == nil {
			t.Fatal("expected error for missing dir, got nil")
		}
	})
	t.Run("Happy", func(t *testing.T) {
		specs := opus47SpecsDir(t)
		stdout, _, err := runCmd(t, "chain", "ExecuteWorkflow", specs)
		if err != nil {
			t.Fatalf("chain happy: unexpected error: %v", err)
		}
		if !strings.Contains(stdout, "Feature Chain:") {
			t.Errorf("expected Feature Chain header, got:\n%s", stdout)
		}
	})
	t.Run("UnknownOp", func(t *testing.T) {
		specs := opus47SpecsDir(t)
		_, _, err := runCmd(t, "chain", "NonExistentOperation99999", specs)
		if err == nil {
			t.Fatal("expected error for unknown operation ID, got nil")
		}
	})
}
