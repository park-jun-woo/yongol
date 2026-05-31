//ff:func feature=cli type=test-helper control=sequence
//ff:what subtestTestGenerateCmdFastAPIBackend — FastAPIBackend 서브테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func subtestTestGenerateCmdFastAPIBackend(t *testing.T) {

	specsDir := filepath.Join(repoRoot(t), "examples", "zenflow", "opus4_7", "specs")
	if _, err := os.Stat(specsDir); err != nil {
		t.Skipf("opus4_7 specs not available: %v", err)
	}
	outDir := t.TempDir()
	_, _, err := runCmd(t, "generate", "--backend", "fastapi", specsDir, outDir)
	_ = err

}
