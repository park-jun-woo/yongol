//ff:func feature=ssac-parse type=test-helper control=iteration dimension=1
//ff:what exerciseFuncDecls — FuncDecl 별 extractParamInfo/collectFuncComments/parseFuncDecl 직접 호출 헬퍼
package ssac

import (
	"go/ast"
	"go/token"
)

// exerciseFuncDecls invokes the per-func ssac parser helpers on each FuncDecl.
func exerciseFuncDecls(fset *token.FileSet, f *ast.File, imports []string, structs []StructInfo) {
	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}
		extractParamInfo(fn)
		collectFuncComments(f, fn.Pos())
		parseFuncDecl(fset, fn, f, "course.ssac", imports, structs)
	}
}
