//ff:func feature=funcspec type=parser control=iteration dimension=1
//ff:what findFuncDeclLine — f.Decls 중 FuncDecl 이름이 일치하는 선언의 줄 번호를 반환
package funcspec

import (
	"go/ast"
	"go/token"
)

// findFuncDeclLine scans f.Decls for a FuncDecl whose name matches
// expectedFuncName and returns its 1-based line number. Returns 0 when no
// matching declaration is found.
func findFuncDeclLine(f *ast.File, fset *token.FileSet, expectedFuncName string) int {
	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Name.Name != expectedFuncName {
			continue
		}
		return fset.Position(fn.Pos()).Line
	}
	return 0
}
