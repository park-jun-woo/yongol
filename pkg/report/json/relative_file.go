//ff:func feature=report type=util control=sequence topic=json
//ff:what relativeFile — specsDir 대비 상대경로 + slash 형식으로 변환 (SARIF 와 동일 로직)
package json

import (
	"path/filepath"
	"strings"
)

// relativeFile rebases a file path against specsDir when possible. Mirrors
// the SARIF emitter behaviour for consistency across formats.
func relativeFile(file, specsDir, absSpecs string) string {
	if file == "" {
		return ""
	}
	if specsDir == "" {
		return filepath.ToSlash(file)
	}
	if rel, err := filepath.Rel(specsDir, file); err == nil && !strings.HasPrefix(rel, "..") {
		return filepath.ToSlash(rel)
	}
	if absSpecs != "" {
		if absFile, err := filepath.Abs(file); err == nil {
			if rel, err := filepath.Rel(absSpecs, absFile); err == nil && !strings.HasPrefix(rel, "..") {
				return filepath.ToSlash(rel)
			}
		}
	}
	return filepath.ToSlash(file)
}
