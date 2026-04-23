//ff:func feature=report type=util control=sequence topic=sarif
//ff:what relativeArtifactURI — file 을 specsDir 기준 상대경로 + slash 형식으로 변환
package sarif

import (
	"path/filepath"
	"strings"
)

// relativeArtifactURI returns file relative to specsDir when possible.
// Falls back to file as-is when either path is empty or not comparable.
func relativeArtifactURI(file, specsDir, absSpecs string) string {
	if file == "" {
		return ""
	}
	if specsDir == "" {
		return filepath.ToSlash(file)
	}
	// Try: file is already relative or under specsDir literally.
	if rel, err := filepath.Rel(specsDir, file); err == nil && !strings.HasPrefix(rel, "..") {
		return filepath.ToSlash(rel)
	}
	// Try: absolute-versus-absolute comparison.
	if absSpecs != "" {
		if absFile, err := filepath.Abs(file); err == nil {
			if rel, err := filepath.Rel(absSpecs, absFile); err == nil && !strings.HasPrefix(rel, "..") {
				return filepath.ToSlash(rel)
			}
		}
	}
	return filepath.ToSlash(file)
}
