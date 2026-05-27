//ff:func feature=cli type=test control=sequence
//ff:what featuresAddCmd test — missing args exit 2

package main

import (
	"testing"
)

func TestFeaturesAddCmd_Missing(t *testing.T) {
	t.Run("Args", func(t *testing.T) {
		_, _, err := runCmd(t, "features", "add")
		if err == nil {
			t.Fatal("expected usage error for missing arg, got nil")
		}
		if !isUsageError(err) {
			t.Fatalf("expected *usageError (exit 2), got %T: %v", err, err)
		}
	})
	t.Run("File", func(t *testing.T) {
		_, _, err := runCmd(t, "features", "add", "/tmp/nonexistent-features.yaml")
		if err == nil {
			t.Fatal("expected error for missing file, got nil")
		}
	})
}
