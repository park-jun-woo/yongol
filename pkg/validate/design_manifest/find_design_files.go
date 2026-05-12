//ff:func feature=validate type=util control=sequence topic=design-manifest
//ff:what findDesignFiles — 지정 디렉토리에서 DESIGN.md 및 *.design.md 파일 검색
package design_manifest

import (
	"path/filepath"
)

// findDesignFiles globs for DESIGN.md and *.design.md under the given directory.
func findDesignFiles(dir string) []string {
	var results []string
	// Pattern 1: DESIGN.md (exact)
	matches, _ := filepath.Glob(filepath.Join(dir, "DESIGN.md"))
	results = append(results, matches...)
	// Pattern 2: *.design.md (convention)
	matches, _ = filepath.Glob(filepath.Join(dir, "*.design.md"))
	results = append(results, matches...)
	return results
}
