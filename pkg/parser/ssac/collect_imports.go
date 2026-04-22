//ff:func feature=ssac-parse type=parser control=iteration dimension=1
//ff:what AST에서 import 경로를 수집
package ssac

import (
	"go/ast"
	"strings"
)

// collectImports는 AST에서 import 경로를 수집한다.
func collectImports(f *ast.File) []string {
	var imports []string
	for _, imp := range f.Imports {
		path := strings.Trim(imp.Path.Value, `"`)
		if path == "net/http" {
			continue
		}
		imports = append(imports, path)
	}
	return imports
}
