//ff:func feature=funcspec type=test control=iteration dimension=1
//ff:what extractReturnTypes / isStubBody / processFuncDecl / findFuncDeclLine / extractGoParseErrorLine
package funcspec

import (
	"go/ast"
)

func firstFunc(f *ast.File) *ast.FuncDecl {
	for _, d := range f.Decls {
		if fn, ok := d.(*ast.FuncDecl); ok {
			return fn
		}
	}
	return nil
}
