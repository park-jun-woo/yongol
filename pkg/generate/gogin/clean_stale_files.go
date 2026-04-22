//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what CleanStaleFiles — 재생성 후 known-set 에 없는 .go 파일을 제거

package gogin

import (
	"os"
	"path/filepath"
	"strings"
)

// CleanStaleFiles removes *.go files in dir whose basename is not present
// in keep. The sweep only touches files that match the ext suffix (e.g.
// ".go", ".hurl") so unrelated artefacts (README.md, openapi.yaml, etc.)
// stay intact. Subdirectories are ignored.
//
// Use this after a WriteManyFiles batch to drop previous-run outputs that
// the new run no longer emits — e.g. when an SSaC function is renamed or
// deleted. Errors while deleting a single file are returned immediately.
func CleanStaleFiles(dir, ext string, keep map[string]bool) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if ext != "" && !strings.HasSuffix(name, ext) {
			continue
		}
		if keep[name] {
			continue
		}
		if err := os.Remove(filepath.Join(dir, name)); err != nil {
			return err
		}
	}
	return nil
}
