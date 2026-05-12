//ff:func feature=validate type=util control=sequence topic=design-manifest
//ff:what normPath — 경로 구분자 정규화 (비교용)
package design_manifest

import (
	"path/filepath"
	"strings"
)

// normPath normalises path separators for comparison.
func normPath(p string) string {
	return strings.ReplaceAll(filepath.Clean(p), "\\", "/")
}
