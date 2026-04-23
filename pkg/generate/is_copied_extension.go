//ff:func feature=generate type=util control=selection
//ff:what isCopiedExtension — React 소스 트리에 복제할 확장자(.tsx/.ts/.jsx/.js/.css/.svg) 화이트리스트
package generate

import (
	"path/filepath"
	"strings"
)

// isCopiedExtension whitelists file types relevant to a React source tree.
// CSS is included because shadcn primitives rely on tailwind utility classes
// but users sometimes override with custom stylesheets.
func isCopiedExtension(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".tsx", ".ts", ".jsx", ".js", ".css", ".svg":
		return true
	}
	return false
}
