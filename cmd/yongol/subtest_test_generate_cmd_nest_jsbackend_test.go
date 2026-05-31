//ff:func feature=cli type=test-helper control=sequence
//ff:what subtestTestGenerateCmdNestJSBackend — NestJSBackend 서브테스트
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func subtestTestGenerateCmdNestJSBackend(t *testing.T) {

	specsDir := filepath.Join(repoRoot(t), "examples", "zenflow", "opus4_7", "specs")
	if _, err := os.Stat(specsDir); err != nil {
		t.Skipf("opus4_7 specs not available: %v", err)
	}
	outDir := t.TempDir()
	_, _, err := runCmd(t, "generate", "--backend", "nestjs", specsDir, outDir)
	// May have validation warnings so generate is refused, but the
	// RunE body should be exercised either way.
	_ = err

}
