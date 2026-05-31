//ff:func feature=cli type=test-helper control=sequence
//ff:what subtestTestGenerateCmdDefaultBackendFromManifest — DefaultBackendFromManifest 서브테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func subtestTestGenerateCmdDefaultBackendFromManifest(t *testing.T) {

	specsDir := filepath.Join(repoRoot(t), "examples", "zenflow", "opus4_7", "specs")
	if _, err := os.Stat(specsDir); err != nil {
		t.Skipf("opus4_7 specs not available: %v", err)
	}
	outDir := t.TempDir()
	// Use empty string for backend to trigger manifest resolution.
	_, _, err := runCmd(t, "generate", "--backend", "", specsDir, outDir)
	_ = err

}
