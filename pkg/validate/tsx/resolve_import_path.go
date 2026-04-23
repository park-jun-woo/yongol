//ff:func feature=validate type=util control=sequence topic=tsx
//ff:what TSX import source 를 절대 경로로 해석 (@/, ./, ../)

package tsx

import (
	"path/filepath"
	"strings"
)

// resolveImportPath converts a TSX import source into an absolute filesystem
// path (without extension). Returns "" when the import is not a local one —
// npm packages should never reach this function since the parser filters
// them, but defensive check remains.
func resolveImportPath(importSource, fromFile, aliasRoot string) string {
	if strings.HasPrefix(importSource, "@/") {
		return filepath.Join(aliasRoot, strings.TrimPrefix(importSource, "@/"))
	}
	if strings.HasPrefix(importSource, "./") || strings.HasPrefix(importSource, "../") {
		return filepath.Join(filepath.Dir(fromFile), importSource)
	}
	return ""
}
