//ff:func feature=cli type=test control=sequence
//ff:what import missing-args — outDir 누락 시 *usageError

package main

import "testing"

// TestIntegrationImport_MissingArgs verifies that `import <source>` (outDir
// missing) yields exit 2 via *usageError — consistent with the other
// subcommands that wrap cobra.ExactArgs through usageArgs.
func TestIntegrationImport_MissingArgs(t *testing.T) {
	_, _, err := runCmd(t, "import", "some-source.yaml")
	if err == nil {
		t.Fatal("expected usage error for missing outDir, got nil")
	}
	if !isUsageError(err) {
		t.Fatalf("expected *usageError (exit 2), got %T: %v", err, err)
	}
}
