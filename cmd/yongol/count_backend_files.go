//ff:func feature=cli type=util control=sequence
//ff:what countBackendFiles — counts .go files under arts/backend (excluding vendor and hidden directories)

package main

import (
	"io/fs"
	"path/filepath"
	"strings"
)

// countBackendFiles returns the number of `.go` files under
// artsDir/backend, ignoring vendor, node_modules, and hidden
// directories. Non-existent directories count as zero so status
// degrades gracefully when the user has not run generate yet.
func countBackendFiles(artsDir string) int {
	root := filepath.Join(artsDir, "backend")
	count := 0
	_ = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() {
			name := d.Name()
			if strings.HasPrefix(name, ".") && path != root {
				return filepath.SkipDir
			}
			if name == "vendor" || name == "node_modules" {
				return filepath.SkipDir
			}
			return nil
		}
		if filepath.Ext(path) == ".go" {
			count++
		}
		return nil
	})
	return count
}
