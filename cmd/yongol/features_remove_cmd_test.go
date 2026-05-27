//ff:func feature=cli type=test control=sequence
//ff:what featuresRemoveCmd test — missing args exit 2, nonexistent specs

package main

import (
	"testing"
)

func TestFeaturesRemoveCmd(t *testing.T) {
	t.Run("MissingArgs", func(t *testing.T) {
		_, _, err := runCmd(t, "features", "remove")
		if err == nil {
			t.Fatal("expected usage error for missing arg, got nil")
		}
		if !isUsageError(err) {
			t.Fatalf("expected *usageError (exit 2), got %T: %v", err, err)
		}
	})
	t.Run("RunE", func(t *testing.T) {
		// Running remove with a valid arg count but missing specs dir hits RunE.
		_, _, err := runCmd(t, "features", "remove", "--yes", "SomeOp")
		if err == nil {
			t.Fatal("expected error for missing specs, got nil")
		}
	})
}
