//ff:func feature=cli type=test control=sequence
//ff:what hashCmd test — missing args, missing dir

package main

import (
	"testing"
)

func TestHashCmd_Missing(t *testing.T) {
	t.Run("Args", func(t *testing.T) {
		_, _, err := runCmd(t, "hash")
		if err == nil {
			t.Fatal("expected usage error for missing arg, got nil")
		}
		if !isUsageError(err) {
			t.Fatalf("expected *usageError (exit 2), got %T: %v", err, err)
		}
	})
	t.Run("Dir", func(t *testing.T) {
		_, _, err := runCmd(t, "hash", "/tmp/nonexistent-yongol-hash")
		if err == nil {
			t.Fatal("expected error for missing dir, got nil")
		}
	})
}
