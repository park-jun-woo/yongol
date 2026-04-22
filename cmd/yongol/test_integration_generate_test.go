//ff:func feature=cli type=test control=iteration dimension=1
//ff:what test: generate 서브커맨드 end-to-end 2 케이스 (args-count / validation-block)

package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestIntegrationGenerate_ArgsCount verifies that `generate <specs>` (artsdir
// missing) exits with code 2 via *usageError, matching validate's
// RangeArgs→usageArgs contract from PhaseC01.
func TestIntegrationGenerate_ArgsCount(t *testing.T) {
	specs := zenflowSpecsDir(t)
	_, _, err := runCmd(t, "generate", specs)
	if err == nil {
		t.Fatal("expected usage error for missing artsDir, got nil")
	}
	if !isUsageError(err) {
		t.Fatalf("expected *usageError (exit 2), got %T: %v", err, err)
	}
}

// TestIntegrationGenerate_ValidationBlock writes a skeletal specs tree
// containing only manifest.yaml. That trips M-1 (model/ has no .go files)
// as an ERROR — generate must refuse with exit 1 and the
// "validation failed: N errors" message.
//
// This exercises the generate gate (PhaseC01): if Validate returns any
// ERROR, printReport yields a non-nil err and Generate is never called.
func TestIntegrationGenerate_ValidationBlock(t *testing.T) {
	tmp := t.TempDir()
	// Minimal manifest.yaml — missing fields still parse; absent model/
	// directory triggers M-1 error.
	manifest := []byte("name: test\nmodule: example.com/test\n")
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
