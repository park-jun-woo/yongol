//ff:func feature=gen-splitter type=util control=sequence
//ff:what importName — ImportSpec 에서 alias 또는 path 마지막 segment 로 접근 식별자 결정
package splitter

import (
	"go/ast"
	"path"
	"strconv"
)

// importName returns the identifier under which an import is referenced
// in source. Priority: explicit alias → last segment of the path, with
// a heuristic that versioned module suffixes (v2, v3, …) fall through to
// their parent segment (foo/v2 → foo), matching Go module semantics.
func importName(imp *ast.ImportSpec) string {
	if imp.Name != nil {
		return imp.Name.Name
	}
	p, err := strconv.Unquote(imp.Path.Value)
	if err != nil {
		return ""
	}
	base := path.Base(p)
	if !isVersionSuffix(base) {
		return base
	}
	return path.Base(path.Dir(p))
}
