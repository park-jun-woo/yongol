//ff:func feature=cli type=test control=sequence
//ff:what initCmd test — missing args, has flags

package main

import (
	"testing"
)

func TestInitCmd(t *testing.T) {
	t.Run("MissingArgs", func(t *testing.T) {
		_, _, err := runCmd(t, "init")
		if err == nil {
			t.Fatal("expected usage error for missing args, got nil")
		}
		if !isUsageError(err) {
			t.Fatalf("expected *usageError (exit 2), got %T: %v", err, err)
		}
	})
	t.Run("HasFlags", func(t *testing.T) {
		cmd := initCmd()
		for _, flag := range []string{"dir", "module", "force"} {
			if cmd.Flags().Lookup(flag) == nil {
				t.Errorf("expected --%s flag", flag)
			}
		}
	})
	t.Run("MissingFeatures", func(t *testing.T) {
		_, _, err := runCmd(t, "init", "TestProject", "/tmp/nonexistent-features.yaml")
		if err == nil {
			t.Fatal("expected error for missing features file, got nil")
		}
	})
	t.Run("WithDescription", func(t *testing.T) {
		_, _, err := runCmd(t, "init", "TestProject", "/tmp/nonexistent-features.yaml", "A test project")
		if err == nil {
			t.Fatal("expected error for missing features file, got nil")
		}
	})
	t.Run("WithDirFlag", func(t *testing.T) {
		_, _, err := runCmd(t, "init", "--dir", t.TempDir(), "TestProject", "/tmp/nonexistent-features.yaml")
		if err == nil {
			t.Fatal("expected error for missing features file, got nil")
		}
	})
}
