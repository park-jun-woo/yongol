//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what writeInitFiles — Python 패키지 __init__.py 파일 일괄 기록

package fastapi

import (
	"os"
	"path/filepath"
)

// writeInitFiles writes __init__.py files for all Python package directories.
func writeInitFiles(appDir string, featureNames []string) error {
	// app/__init__.py
	if err := os.WriteFile(filepath.Join(appDir, "__init__.py"), []byte(""), 0o644); err != nil {
		return err
	}

	// Ensure services and routers __init__.py exist
	dirs := []string{"services", "routers", "schemas"}
	for _, d := range dirs {
		dir := filepath.Join(appDir, d)
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return err
		}
		if err := os.WriteFile(filepath.Join(dir, "__init__.py"), []byte(""), 0o644); err != nil {
			return err
		}
	}

	return nil
}
