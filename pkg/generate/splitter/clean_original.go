//ff:func feature=gen-splitter type=util control=iteration dimension=1
//ff:what cleanOriginal — 분할 완료 후 원본 파일(server.gen.go / models.go / *.sql.go)을 삭제
package splitter

import (
	"os"
	"path/filepath"
)

// cleanOriginal deletes the unsplit source files that the external
// generator emitted into dir. The deletion criteria per tool:
//
//	oapi-codegen → remove any *.gen.go that sits alongside split
//	               outputs (callers pass keep to exclude split results)
//	sqlc         → remove models.go and *.sql.go whose basename matches
//	               one of the original query files (keep list = split
//	               outputs + preserved_files.go entries)
//
// keep holds file names (basenames) that must survive. Errors while
// deleting a single file are returned immediately.
func cleanOriginal(dir string, tool Tool, keep map[string]bool) error {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if keep[name] {
			continue
		}
		if isPreservedFile(name, tool) {
			continue
		}
		if !matchesOriginal(name, tool) {
			continue
		}
		if err := os.Remove(filepath.Join(dir, name)); err != nil {
			return err
		}
	}
	return nil
}
