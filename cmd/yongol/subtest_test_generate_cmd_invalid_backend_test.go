//ff:func feature=cli type=test-helper control=sequence
//ff:what subtestTestGenerateCmdInvalidBackend — InvalidBackend 서브테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func subtestTestGenerateCmdInvalidBackend(t *testing.T) {

	specsDir := filepath.Join(repoRoot(t), "examples", "zenflow", "opus4_7", "specs")
	if _, err := os.Stat(specsDir); err != nil {
		t.Skipf("opus4_7 specs not available: %v", err)
	}
	outDir := t.TempDir()
	_, _, err := runCmd(t, "generate", "--backend", "invalid-framework", specsDir, outDir)
	// The generate may fail during validate or at backend resolution, both are acceptable.
	if err == nil {
		t.Fatal("expected error for invalid backend, got nil")
	}

}
