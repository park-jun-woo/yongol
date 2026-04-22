//ff:func feature=gen-splitter type=command control=iteration dimension=1
//ff:what SplitDirectory — 디렉토리의 도구 산출 파일을 모두 분할 + 원본 정리 (일괄 처리 엔트리)
package splitter

import (
	"fmt"
	"os"
	"path/filepath"
)

// SplitDirectory is the top-level entry used by run_backend.go. It scans
// dir for files matching tool's source pattern, splits each into
// per-decl files, and finally removes the originals (while keeping
// preserved files like sqlc's querier.go/db.go). If dir does not exist
// the call is a no-op — tools that were not run produce no directory.
func SplitDirectory(dir string, tool Tool) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return nil
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	keep := map[string]bool{}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if isPreservedFile(name, tool) {
			keep[name] = true
			continue
		}
		full := filepath.Join(dir, name)
		if !isSourceFile(full, tool) {
			continue
		}
		written, err := SplitFile(full, dir, tool)
		if err != nil {
			return fmt.Errorf("split %s: %w", name, err)
		}
		for _, w := range written {
			keep[w] = true
		}
	}
	return cleanOriginal(dir, tool, keep)
}
