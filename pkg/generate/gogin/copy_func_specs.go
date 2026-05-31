//ff:func feature=gen-gogin type=util control=sequence
//ff:what copyFuncSpecs — specs/func/ 하위 Go 파일을 arts/backend/internal/ 에 복사
package gogin

import (
	"fmt"
	"os"
	"path/filepath"
)

// copyFuncSpecs copies specs/func/<pkg>/*.go → arts/backend/internal/<pkg>/*.go.
// File contents are preserved as-is (no transformation).
// If specs/func/ does not exist, the step is silently skipped.
func copyFuncSpecs(specsDir, artifactsDir string) error {
	funcDir := filepath.Join(specsDir, "func")
	if _, err := os.Stat(funcDir); os.IsNotExist(err) {
		return nil // skip — no func specs
	}

	return filepath.Walk(funcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(funcDir, path)
		if err != nil {
			return fmt.Errorf("rel path: %w", err)
		}
		dst := filepath.Join(artifactsDir, "backend", "internal", rel)

		if info.IsDir() {
			return os.MkdirAll(dst, 0o755)
		}
		// only copy .go files
		if filepath.Ext(path) != ".go" {
			return nil
		}
		return copyFile(path, dst)
	})
}
