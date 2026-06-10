//ff:func feature=cli type=test control=sequence
//ff:what generate validation-block — ERROR 있으면 generate 거절 (exit 1)

package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestIntegrationGenerate_ValidationBlock writes a skeletal specs tree
// containing only manifest.yaml. That trips M-1 (model/ has no .go files)
// as an ERROR — generate must refuse with exit 1 and the
// "validation failed: N errors" message.
//
// This exercises the generate gate (PhaseC01): if Validate returns any
// ERROR, printReport yields a non-nil err and Generate is never called.
func TestIntegrationGenerate_ValidationBlock(t *testing.T) {
	tmp := t.TempDir()
	// Minimal, schema-valid manifest.yaml — it parses cleanly (no unknown
	// keys, so strict decoding accepts it); the absent model/ directory then
	// triggers M-1 as a validation ERROR, which is what this test exercises.
	manifest := []byte("apiVersion: yongol/v1\nkind: Project\nmetadata:\n  name: test\nbackend:\n  module: example.com/test\n")
	if err := os.WriteFile(filepath.Join(tmp, "manifest.yaml"), manifest, 0644); err != nil {
		t.Fatalf("write manifest: %v", err)
	}

	stdout, _, err := runCmd(t, "generate", tmp, filepath.Join(tmp, "arts"))
	if err == nil {
		t.Fatalf("expected validation error, got nil\nstdout:\n%s", stdout)
	}
	if isUsageError(err) {
		t.Fatalf("validation failure must be exit 1, not usage error: %v", err)
	}
	if !strings.Contains(err.Error(), "validation failed") &&
		!strings.Contains(err.Error(), "generate refused") {
		t.Errorf("expected error to mention `validation failed` or `generate refused`, got: %v", err)
	}
	if !strings.Contains(stdout, "errors") {
		t.Errorf("expected stdout to show error count line, got:\n%s", stdout)
	}
}
