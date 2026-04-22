//ff:func feature=ssac-parse type=parser control=iteration dimension=1
//ff:what AST에서 struct 선언을 수집
package ssac

import "go/ast"

// collectStructs는 AST에서 struct 선언을 수집한다.
func collectStructs(f *ast.File) []StructInfo {
	var structs []StructInfo
	for _, decl := range f.Decls {
		structs = append(structs, collectStructsFromDecl(decl)...)
	}
	return structs
}
